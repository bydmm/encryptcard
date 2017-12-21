package main

import (
	"bufio"
	"encryptcard/share"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/davyxu/golog"
)

var log = golog.New("main")
var blockCount = 0

// 清空屏幕
func clearScreen() {
	fmt.Printf("\033[2J\033[0;0H")
}

// 隐藏光标
func hideCursor() {
	fmt.Printf("\033[?25l")
}

// 显示光标
func showCursor() {
	fmt.Printf("\033[?25h")
}

// 让mac说一句话
func say(word string) {
	cmd := exec.Command("say", word)
	if err := cmd.Run(); err != nil {

	}
}

// 抽卡动画
func animation() {
	clearScreen()
	hideCursor()
	assets := AssetNames()
	sort.Strings(assets)
	for _, file := range assets {
		raw, err := Asset(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s", raw)
		time.Sleep(10000000)
		fmt.Printf("\033[0;0H")
	}
	showCursor()
}

// 初始化游戏
func initGame(concurrency int) {
	// 初始化种子
	rand.Seed(time.Now().UnixNano())
	// 创建文件夹
	os.Mkdir("./saves", 0755)
	// concurrency
	runtime.GOMAXPROCS(concurrency)
	// 屏蔽socket层的调试日志
	golog.SetLevelByString("cellnet", "error")
	golog.SetLevelByString("main", "error")
}

// 开始屏幕
func startScreen(sound bool, concurrency int) {
	clearScreen()
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n")
	fmt.Printf("*                                                                                             *\n")
	fmt.Printf("*         __________________                                                                  *\n")
	fmt.Printf("*        |     ☆ ☆ ★ ☆ ☆    |                                                                 *\n")
	fmt.Printf("*        |                  |                                                                 *\n")
	fmt.Printf("*        |     1 1  2 2 2   |          EncryptCard                                            *\n")
	fmt.Printf("*        |    1  1     2    |                                                                 *\n")
	fmt.Printf("*        |   1 1 1    2     |          v0.0.1                                                 *\n")
	fmt.Printf("*        |       1   2      |                                                                 *\n")
	fmt.Printf("*        |       1  2 2 2   |          Use Proof-of-Work digging card                         *\n")
	fmt.Printf("*        |                  |          with Distributed Game Archives System                  *\n")
	fmt.Printf("*        |       Answer     |                                                                 *\n")
	fmt.Printf("*        |                  |          Project: https://github.com/bydmm/encryptcard          *\n")
	fmt.Printf("*        |   A T K   D E F  |                                                                 *\n")
	fmt.Printf("*        |    9 9     9 9   |                                                                 *\n")
	fmt.Printf("*        |__________________|          Concurrency: %d                                         *\n", concurrency)
	fmt.Printf("*                                                                                             *\n")
	fmt.Printf("*                                                                                             *\n")
	fmt.Printf("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *\n")
	time.Sleep(1 * time.Second)
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

// func cardBlockDig(card share.CardBlock) cardproto.CardBlockDig {
// 	return cardproto.CardBlockDig{
// 		Version:    card.Version,
// 		Hard:       card.Hard,
// 		PubKey:     card.PubKey,
// 		Timestamp:  card.Timestamp,
// 		RandNumber: card.RandNumber,
// 		PrevCardID: card.PrevCardID,
// 	}
// }
// func receiveDcardBlock(msg *cardproto.CardBlockReceive) share.CardBlock {
// 	return share.CardBlock{
// 		Version:    msg.Version,
// 		Hard:       msg.Hard,
// 		PubKey:     msg.PubKey,
// 		Timestamp:  msg.Timestamp,
// 		RandNumber: msg.RandNumber,
// 		PrevCardID: msg.PrevCardID,
// 	}
// }

// 开始挖卡
func start(sound bool, concurrency int) {
	startScreen(sound, concurrency)
	// 初始化
	initGame(concurrency)

	// 用户钥匙对
	key := share.GetKeyPair()
	userPubKey := share.PubKey(key)
	headerBlock := share.CardBlock{}

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

	// cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {
	// 	log.Debugln("client connected")
	// 	// ev.Ses.Send(&cardproto.CardBlockFetch{})
	// })

	// 多开几个挖卡的任务
	for index := 0; index < concurrency; index++ {
		go func() {
			for {
				// 使用算力工作证明无限抽卡
				block := share.CardBlock{PubKey: userPubKey, Hard: 0, PrevCardID: headerBlock.CardID(), Height: 0}
				CardID := block.Build()
				if CardID != "" {
					card, _ := block.Card()
					// fmt.Printf("%d-%d-%d\n", card.ID, card.Attack, card.Defense)
					if card.ID == 0 && card.Attack == 99 && card.Defense == 99 {
						saveBlock(block)
					}
					// msg := cardBlockDig(block)
					// ev.Ses.Send(&msg)
				}
				blockCount++
			}
		}()
	}

	// 定时显示挖卡的速度
	go func() {
		second := 0
		for {
			time.Sleep(1 * time.Second)
			second++
			speed := blockCount / second
			fmt.Printf("%d Blocks in %d second (%d/s)\r", blockCount, second, speed)
		}
	}()

	// queue.StartLoop()

	// queue.Wait()

	select {}
}

func main() {
	concurrency := flag.Int("c", 1, "并发数，默认为1, 不建议超过CPU数")
	flag.Parse()

	sound := openSound()
	start(sound, *concurrency)
}
