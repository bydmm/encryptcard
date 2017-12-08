package main

import (
	"fmt"
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
			fmt.Println(json)
		}
		// 交出控制权，不然卡死cpu了。
		time.Sleep(1)
	}
}
