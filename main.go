package main

import (
	"bufio"
	"crypto/rsa"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strings"
	"time"
)

func clearScreen() {
	fmt.Printf("\033[2J\033[0;0H")
}

func hideCursor() {
	fmt.Printf("\033[?25l")
}

func showCursor() {
	fmt.Printf("\033[?25h")
}

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

func initGame() {
	// 初始化种子
	rand.Seed(time.Now().UnixNano())
	// 创建文件夹
	os.Mkdir("./saves", 0755)
}

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
}

func openSound() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("[此功能仅对OSX有效]开启声音？ 输入[Y]确认: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	return strings.ToUpper(text) == "Y"
}

func whenFindCard(key *rsa.PrivateKey, block CardBlock, sound bool) {
	block.Signature = block.sign(key)
	card, err := block.card()
	if err == nil {
		animation()
		clearScreen()
		fmt.Printf("id: %d\n", card.id)
		fmt.Printf("attack: %d\n", card.attack)
		fmt.Printf("defense: %d\n", card.defense)
		path := fmt.Sprintf("./saves/%d_%d_%d.json", card.id, card.attack, card.defense)
		fmt.Printf("save to :%s\n", path)
		f, err := os.Create(path)
		f.WriteString(block.json())
		if err != nil {
			panic(err)
		}
		c, ok := CardPrototypes[card.id]
		if ok {
			fmt.Printf("%s: %s\n", c.name, c.Lines)
			if sound {
				say(c.Lines)
			}
		}
	}
}

func digging(key *rsa.PrivateKey, user string, sound bool) {
	block := CardBlock{PubKey: user}
	block.build()
	if block.CardID != "" {
		whenFindCard(key, block, sound)
	}
}

func start(sound bool, concurrency int) {
	startScreen(sound, concurrency)
	initGame()
	time.Sleep(2000000000)
	// 用户钥匙对
	key := getKeyPair()
	user := pubKey(key)
	// start := time.Now().UnixNano()
	runtime.GOMAXPROCS(concurrency)
	// chann := make(chan int, 100000000)
	for index := 0; index < concurrency; index++ {
		go func() {
			for {
				// 使用算力工作证明无限抽卡
				digging(key, user, sound)
				// chann <- 1
				// speed := int(time.Now().UnixNano()-start) / len(chann)
				// fmt.Printf("%d Block per second\n", speed)
			}
		}()
	}
	select {}
}

func verifyCard(verifyPath *string) {
	// 验证card文件
	fmt.Printf("校验: %s\n", *verifyPath)
	card := loadCard(*verifyPath)
	if card.verify() {
		fmt.Printf("校验成功\n")
	}
}

func main() {
	verifyPath := flag.String("v", "", "验证卡片json文件")
	concurrency := flag.Int("c", 1, "并发数，默认为1, 不建议超过CPU数")
	flag.Parse()
	if *verifyPath != "" {
		verifyCard(verifyPath)
		os.Exit(1)
	}

	sound := openSound()
	start(sound, *concurrency)
}
