package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 用户公钥
	user := "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2"
	// 使用算力工作证明无限抽卡
	for true {
		block := CardBlock{PubKey: user}
		json := block.build()
		if json != "" {
			card, err := block.card()
			if err == nil {
				fmt.Printf("---------------------------\n")
				fmt.Printf("id: %d\n", card.id)
				fmt.Printf("attack: %d\n", card.attack)
				fmt.Printf("defense: %d\n", card.defense)
				path := fmt.Sprintf("./saves/%d_%d_%d.json", card.id, card.attack, card.defense)
				fmt.Printf("save to :%s\n", path)
				f, err := os.Create(path)
				f.WriteString(json)
				if err != nil {
					panic(err)
				}
			}
		}
		// 交出控制权，不然卡死cpu了。
		time.Sleep(1)
	}
}
