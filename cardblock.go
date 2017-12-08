package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

// CardBlock is a good block
// {
// 	"pubkey": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2",
// 	"timestamp": 1974545345345,
// 	"randNumber": 6653,
// 	"cardBlock": "15c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800000009004",
// 	"signature": "dsfsdf34515c9c6c3afb2b2ff612c5ea37b563c50dac4e95d7a93695bc5d6800",
// 	"owner": "1ccfce1ed647ec3b12c398f4791a1adb3285cfff85ce7d382362c321a1a1df2"
//   }
type CardBlock struct {
	PubKey     string
	Timestamp  string
	RandNumber string
	CardID     string
	Signature  string
}

// 时间戳
func timestamp() string {
	timestamp := time.Now().UnixNano()
	return strconv.Itoa(int(timestamp))
}

// 随机数
func randNumber() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return strconv.Itoa(r.Intn(1000))
}

// 判定是否为卡的函数
func findCard(hashCard [32]byte) string {
	// 难度系数
	hard := 4
	// 获取卡的hash值
	card := hex.EncodeToString(hashCard[:])
	// 截取这个卡的最后几位
	last := string(card[len(card)-(hard+7):])

	// 难度系数就是说，最后几位的开头要有几个0
	// 由于这个hash应该是随机分布的，那么0越多自然越难
	headZero := ""
	for index := 0; index < hard; index++ {
		headZero += "0"
	}
	if last[0:hard] != headZero {
		return ""
	}

	// 满足hard个0后，还要是个数字，否则匹配不到卡的id
	i, err := strconv.ParseInt(last, 10, 32)
	if err != nil || i == 0 {
		return ""
	}

	return card
}

func (card CardBlock) build() string {
	card.Timestamp = timestamp()
	card.RandNumber = randNumber()
	// 使用用户公钥，时间戳以及随机数作为种子
	key := card.PubKey + card.Timestamp + card.RandNumber
	// 去生成一个hash值，这里使用sha256这个比较公允的算法
	rawOre := sha256.Sum256([]byte(key))
	// 根据规则去判断hash是否是一张卡
	cardHash := findCard(rawOre)

	if cardHash != "" {
		card.CardID = cardHash
		json, _ := json.Marshal(card)
		return string(json)
	}
	return ""
}
