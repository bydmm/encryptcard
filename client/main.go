package main

import (
	"bufio"
	"crypto/rsa"
	"encryptcard/protoc/cardproto"
	"encryptcard/share"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
)

// ChainPath 区块链文件路径
const ChainPath = "./blocks/cards.chain"

// 初始化log
var log = golog.New("main")

// 初始化挖卡计数
var blockCount = 0

// cardblockChain 内存中的区块链
var cardblockChain = &share.CardBlockChain{}

// userKeyPair 用户密钥对
var userKeyPair *rsa.PrivateKey

// userPubKey 用户公钥
var userPubKey string

// 并发数
var maxConcurrency = 1

// 启动声音
var enableSound = false

// 启动动画
var enableAnimation = false

// 挖矿开关
var doneSync = make(chan bool, 1)

// 动画正在播放的flag
var playingSync = make(chan bool, 1)

// 初始化区块链
func initChain() {
	log.Infof("从磁盘读取区块链....\n")
	if _, err := os.Stat(ChainPath); os.IsNotExist(err) {
		log.Infof("区块文件不存在....\n")
	} else {
		// 如果文件存在,试图读取区块
		cardblockChain = share.LoadCardBlockChainFromDisk(ChainPath)
		log.Infof("区块链读取完成....\n")
		log.Infof("区块高度: %d\n", len(cardblockChain.Cardblocks))
	}
}

// 初始化游戏
func initGame() {
	// 初始化种子
	rand.Seed(time.Now().UnixNano())

	// 创建文件夹
	os.Mkdir("./saves", 0755)
	os.Mkdir("./blocks", 0755)

	// 最大并发
	runtime.GOMAXPROCS(maxConcurrency)

	// 屏蔽socket层的调试日志
	golog.SetLevelByString("cellnet", "error")
	golog.SetLevelByString("main", "info")

	// 初始化区块
	initChain()

	// 初始化密钥
	userKeyPair = share.GetKeyPair()
	userPubKey = share.PubKey(userKeyPair)

	// 初始化动画同步
	playingSync <- true
}

// 启用动画效果
func openAnimation() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n抽卡动画仅在osx, linux, Win10 Powershell中可用, 并且可能降低挖卡效率\n")
	fmt.Print("开启抽卡动画？ 输入[Y]确认: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return strings.ToUpper(text) == "Y"
}

// 打开声音
func openSound() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("开启声音[此功能仅对OSX有效] 输入[Y]确认: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return strings.ToUpper(text) == "Y"
}

// 当挖到卡后
func whenFindCard(block *share.CardBlock) {
	<-playingSync
	card, err := block.Card()
	if err == nil {
		if enableAnimation {
			animation()
			clearScreen()
		}
		c, ok := share.CardPrototypes[card.ID]
		if ok {
			fmt.Printf("%s: %s 攻击: %d 防御: %d\n", c.Name, c.Lines, card.Attack, card.Defense)
			if enableSound {
				say(c.Lines)
			}
		}
	}
	playingSync <- true
}

// SaveChainToDisk 保存区块到磁盘
func SaveChainToDisk() {
	share.Store(cardblockChain, ChainPath)
}

// ReSyncChain 区块被破坏,重新同步
func ReSyncChain(peer cellnet.Peer) {
	cardblockChain = &share.CardBlockChain{}
	CheckSync(peer)
}

// CheckSync 同步检查
func CheckSync(peer cellnet.Peer) {
	if cardblockChain.Cardblocks == nil {
		log.Infof("本地无区块缓存...从创始区块开始同步\n")
		RequestCardBlocksFetch(peer, 0)
	} else {
		height := cardblockChain.Height()
		log.Infof("本地高度%d, 请求与服务器同步\n", height)
		RequestCardBlockSync(peer, cardblockChain.Height(), cardblockChain.HeadBlock().CardID())
	}

}

// SyncResponse 获取同步检查的结果
func SyncResponse(peer cellnet.Peer) {
	// 检查与服务器同步的状态
	CardBlockSyncResponse(peer, func(msg *cardproto.CardBlockSyncResponse) {
		height := cardblockChain.Height()
		if msg.Valid {
			headBlock := msg.CardBlock
			if headBlock.Height > height {
				RequestCardBlocksFetch(peer, height+1)
			} else {
				// 开挖, 解除阻塞,开启挖矿
				doneSync <- true
				log.Infof("与主链同步完成,当前高度: %d, 难度: %d, 目标难度: %d\n", height, headBlock.Hard, cardblockChain.AdaptiveHard())
			}
		} else {
			log.Errorln("与主链失去同步, 正在清除本地缓存重新同步...\n")
			ReSyncChain(peer)
		}
	})
}

// FetchBlocks 从服务器获取大量区块
func FetchBlocks(peer cellnet.Peer) {
	CardBlockFetchResponse(peer, func(msg *cardproto.CardBlockFetchResponse) {
		if msg.Valid {
			log.Debugf("msg.CardBlocks count: %d", len(msg.CardBlocks))
			// 确认还在链上
			for index := 0; index < len(msg.CardBlocks); index++ {
				remoteBlock := share.ProtoToBlock(msg.CardBlocks[index])

				log.Debugf("remoteBlock.Height: %d", remoteBlock.Height)

				// 处理创世区块同步的问题
				if remoteBlock.Height == 0 {
					cardblockChain.Cardblocks = []*share.CardBlock{remoteBlock}
					continue
				}

				if remoteBlock.PrevCardID == cardblockChain.HeadBlock().CardID() {
					cardblockChain.Cardblocks = append(cardblockChain.Cardblocks, remoteBlock)
				} else {
					log.Fatalf("似乎出现了线程不安全的情况...\n")
				}
			}

			height := cardblockChain.Height()
			if msg.Finish {
				CheckSync(peer)
			} else {
				log.Infof("同步未完成..当前高度%d\n", height)
				RequestCardBlocksFetch(peer, height+1)
			}

			SaveChainToDisk()
		} else {
			log.Fatalf("与主链失去同步, 建议删除本地区块再试...\n")
		}
	})
}

// CardBlockMsg 从服务器获得最新区块推送
func CardBlockMsg(peer cellnet.Peer) {
	// 从服务器得知最新区块
	CardBlockLiveMsg(peer, func(msg *cardproto.CardBlock) {
		block := share.ProtoToBlock(msg)
		//判断是否能接受这张卡
		if block.VerifyCardID() == false {
			log.Errorf("无法验证的区块推送: %s", block.CardID())
			return
		}
		// 如果本地区块跟不上了
		if block.Height > (cardblockChain.Height() + 1) {
			// 启动同步检查
			CheckSync(peer)
			return
		}

		// 添加到本地块
		cardblockChain.Cardblocks = append(cardblockChain.Cardblocks, block)

		if block.PubKey == userPubKey {
			go whenFindCard(block)
		}

		log.Infof("同步新区块,当前区块高度: %d 当前区块难度: %d 目标难度:%d\n", block.Height, block.Hard, cardblockChain.AdaptiveHard())
	})
}

// StoreChainToDisk

// SyncBlocks 与服务器同步区块
func SyncBlocks(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event) {
	// 启动同步检查
	CheckSync(peer)

	// 从服务器得知同步结果
	SyncResponse(peer)

	// 从服务器获取区块
	FetchBlocks(peer)

	// 接受服务器区块推送
	CardBlockMsg(peer)
}

// Miner 辛苦的矿工,挖挖挖
func Miner(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event) {
	<-doneSync
	//等待挖矿完成
	log.Infof("开始挖卡...\n")
	// 根据配置启动多个挖卡够程(goroutine)
	for index := 0; index < maxConcurrency; index++ {
		go func() {
			for {
				// 使用算力工作证明无限抽卡
				headBlock := cardblockChain.HeadBlock()
				block := share.CardBlock{PubKey: userPubKey, Hard: cardblockChain.AdaptiveHard(), PrevCardID: headBlock.CardID(), Height: (headBlock.Height + 1)}
				CardID := block.Build()
				if CardID != "" {
					if block.VerifyCardID() {
						RequestCardBlockPush(peer, &block)
					}
				}
				blockCount++
			}
		}()
	}

	// 定时显示挖卡的速度
	go func() {
		second := 0
		for {
			<-playingSync
			time.Sleep(1 * time.Second)
			second++
			speed := blockCount / second
			fmt.Printf("%d Blocks in %d second (%d/s)\r", blockCount, second, speed)
			playingSync <- true
		}
	}()
}

// AfterConnect 当连接到服务器之后
func AfterConnect(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event) {
	//与服务器同步区块
	SyncBlocks(queue, peer, ev)

	//挖矿子程序
	go Miner(queue, peer, ev)
}

// Start 开始主程序
func Start() {
	// 开始链接服务器
	log.Infof("与服务器建立链接.....\n")
	StartClient(func(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event, success bool) {
		if success {
			log.Infof("成功与服务器建立链接...\n")
			AfterConnect(queue, peer, ev)
		} else {
			log.Infof("与服务器断开...\n")
		}
	})
}

// MyCards 显示我的卡组
func MyCards() {
	if cardblockChain.Cardblocks == nil {
		panic("本地无可用的区块文件，无法查询")
	}

	blocks := []*share.CardBlock{}
	for index := 0; index < len(cardblockChain.Cardblocks); index++ {
		block := cardblockChain.Cardblocks[index]
		if block.PubKey == userPubKey {
			whenFindCard(block)
			blocks = append(blocks, block)
		}
	}
	if len(blocks) > 0 {
		fmt.Printf("恭喜你，总计%d张卡\n", len(blocks))
	} else {
		fmt.Printf("别气馁，努力挖挖总是有的\n")
	}

}

func main() {
	my := flag.Bool("my", false, "查看我拥有的卡")
	concurrency := flag.Int("c", 1, "并发数,默认为1, 不建议超过CPU数")

	flag.Parse()

	// 并发限制
	maxConcurrency = *concurrency
	// 声音限制
	enableSound = openSound()
	// 声音限制
	enableAnimation = openAnimation()

	// 启动画面
	startScreen()

	// 初始化
	initGame()

	if *my {
		MyCards()
	} else {
		Start()
	}
}
