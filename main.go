package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 抽卡动画
func animation() {
	files, err := filepath.Glob("./animation/*.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		raw, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("\033[0;0H")
		fmt.Printf("%s", raw)
		time.Sleep(10000)
	}
}

func initGame() {
	fmt.Printf("---------------------------\n")
	os.Mkdir("./saves", 644)
}

func start() {
	initGame()
	// 用户钥匙对
	key := getKeyPair()

	user := pubKey(key)
	// 使用算力工作证明无限抽卡
	for true {
		block := CardBlock{PubKey: user}
		block.build()
		if block.CardID != "" {
			block.Signature = block.sign(key)
			card, err := block.card()
			if err == nil {
				animation()
				fmt.Printf("---------------------------\n")
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
			}
		}
		// 交出控制权，不然卡死cpu了。
		time.Sleep(1)
	}
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
	flag.Parse()
	if *verifyPath != "" {
		verifyCard(verifyPath)
		os.Exit(1)
	}

	start()
}
