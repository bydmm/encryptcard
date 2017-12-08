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

// 用户挖到卡后，展示自己的钥匙，当时挖出的时间戳，以及这个随机数
// 任何人都可以还原这个算法，证明这张卡确实是这个用户挖出来的
// 用户之间也许还可以通过私钥来达成交易(未实现)

// {
// 	"pubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2",
// 	"timestamp": 1974545345345,
// 	"randNumber": 6653,
// 	"cardBlock": "15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800000009004",
// 	"signature": "dsfsdf34515c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800",
// 	"owner": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2"
//   }

//   通过这个验证卡是正确的（不挖找不到这个block）
//   cardBlock = sha256(pubkey + timestamp + randNumber + cardBlock)

//   通过这个来验证你拥有这个块，用来交易
//   signature = 签名函数(private_key, cardBlock)
//   cardBlock = 验证函数(pubkey, signature)

//   cardBlock

//   15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d68(00000000019999)

//   0000000 基础难度系数，0越多总体的难度提升
//   001 卡id，设定为0越多卡越稀有
//   99  攻击，纯属娱乐
//   99  防御，纯属娱乐

//   交易出去的卡：
//   {
// 	"pubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2",
// 	"timestamp": 1974545345345,
// 	"randNumber": 6653,
// 	"cardBlock": "15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800000009004",
// 	"signature": "dsfsdf34515c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800",
// 	"ownerPubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2"
//   }

//   signature = 签名函数(创造者的private_key, (ownerPubkey + cardBlock))
//   ownerPubkey + cardBlock = 验证函数(pubkey, signature)
