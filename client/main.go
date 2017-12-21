package main

import (
	"bufio"
	"crypto/rsa"
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
var cardblockChain *share.CardBlockChain

// userKeyPair 用户密钥对
var userKeyPair *rsa.PrivateKey

// 并发数
var maxConcurrency int

// 启动声音
var enableSound bool

// 初始化区块链
func initChain() {
	log.Infof("从磁盘读取区块链....\n")
	if _, err := os.Stat(ChainPath); os.IsNotExist(err) {
		log.Infof("区块文件不存在....\n")
	} else {
		// 如果文件存在，试图读取区块
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
	golog.SetLevelByString("main", "error")

	// 初始化区块
	initChain()

	// 初始化密钥
	userKeyPair = share.GetKeyPair()
}

// 打开声音
func openSound() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[此功能仅对OSX有效]开启声音？ 输入[Y]确认: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return strings.ToUpper(text) == "Y"
}

func saveBlock(block share.CardBlock) {
	path := fmt.Sprintf("./saves/%s.json", block.CardID())
	fmt.Printf("save to :%s\n", path)
	f, err := os.Create(path)
	f.WriteString(block.JSON())
	if err != nil {
		panic(err)
	}
}

func showCardInfo(block share.CardBlock) {
	card, err := block.Card()
	if err != nil {
		log.Infof("error card: %d\n", block.CardID())
	}
	log.Infof("CardBlockReceive: %s\n", block.CardID())
	log.Infof("id: %d\n", card.ID)
	log.Infof("attack: %d\n", card.Attack)
	log.Infof("defense: %d\n", card.Defense)
}

// 当挖到卡后
func whenFindCard(block share.CardBlock, sound bool) {
	card, err := block.Card()
	if err == nil {
		// animation()
		// clearScreen()
		c, ok := share.CardPrototypes[card.ID]
		if ok {
			fmt.Printf("%s: %s\n", c.Name, c.Lines)
			if sound {
				say(c.Lines)
			}
		}
	}
}

// Start 开始主程序
func Start() {
	startScreen()
	// 初始化
	initGame()

	// 开始链接服务器
	log.Infof("与服务器建立链接.....\n")
	StartClient(func(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event, success bool) {
		if success {
			log.Infof("成功与服务器建立链接...\n")
		} else {
			log.Errorf("与服务器断开...\n")
		}

	})

	// cellnet.RegisterMessage(peer, "cardproto.CardBlockReceive", func(ev *cellnet.Event) {
	// 	msg := ev.Msg.(*cardproto.CardBlockReceive)
	// 	block := receiveDcardBlock(msg)
	// 	if block.Verify() == false {
	// 		log.Infof("假卡！？\n")
	// 		return
	// 	}
	// 	headerBlock = block
	// 	saveBlock(headerBlock)
	// 	showCardInfo(headerBlock)

	// 	if headerBlock.PubKey == userPubKey {
	// 		whenFindCard(block, sound)
	// 	}
	// })

	// 多开几个挖卡的任务
	// for index := 0; index < maxConcurrency; index++ {
	// 	go func() {
	// 		for {
	// 			// 使用算力工作证明无限抽卡
	// 			userPubKey := share.PubKey(userKeyPair)
	// 			block := share.CardBlock{PubKey: userPubKey, Hard: 0, PrevCardID: cardblockChain.HeadBlock().CardID(), Height: 0}
	// 			CardID := block.Build()
	// 			if CardID != "" {
	// 				card, _ := block.Card()
	// 				// fmt.Printf("%d-%d-%d\n", card.ID, card.Attack, card.Defense)
	// 				if card.ID == 0 && card.Attack == 99 && card.Defense == 99 {
	// 					saveBlock(block)
	// 				}
	// 				// msg := cardBlockDig(block)
	// 				// ev.Ses.Send(&msg)
	// 			}
	// 			blockCount++
	// 		}
	// 	}()
	// }

	// // 定时显示挖卡的速度
	// go func() {
	// 	second := 0
	// 	for {
	// 		time.Sleep(1 * time.Second)
	// 		second++
	// 		speed := blockCount / second
	// 		fmt.Printf("%d Blocks in %d second (%d/s)\r", blockCount, second, speed)
	// 	}
	// }()

	// queue.StartLoop()

	// queue.Wait()

	// select {}
}

func main() {
	concurrency := flag.Int("c", 1, "并发数，默认为1, 不建议超过CPU数")
	flag.Parse()

	// 并发现只
	maxConcurrency = *concurrency
	// 声音限制
	enableSound = openSound()

	Start()
}
