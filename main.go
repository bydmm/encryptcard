package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func start() {
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

func main() {
	verifyPath := flag.String("v", "", "验证卡片json文件")
	flag.Parse()

	// 验证card文件
	if *verifyPath != "" {
		fmt.Printf("校验: %s\n", *verifyPath)
		card := loadCard(*verifyPath)
		if card.verify() {
			fmt.Printf("校验成功\n")
		}
		os.Exit(1)
	}

	start()
}
