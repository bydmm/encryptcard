package main

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 难度系数
const hard int = 4
const version = "v0.0.1"

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
	Version    string
	PubKey     string
	Timestamp  string
	RandNumber string
	Hard       string
	CardID     string
	Signature  string
}

// Card is CardBlock helper
type Card struct {
	id      int
	attack  int
	defense int
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

func (card *CardBlock) sign(key *rsa.PrivateKey) string {
	m := card.Version + card.PubKey + card.Timestamp + card.RandNumber +
		card.Hard + card.CardID
	message := []byte(m)

	// Only small messages can be signed directly; thus the hash of a
	// message, rather than the message itself, is signed. This requires
	// that the hash function be collision resistant. SHA-256 is the
	// least-strong hash function that should be used for this at the time
	// of writing (2016).
	hashed := sha256.Sum256(message)

	signature, err := rsa.SignPKCS1v15(crand.Reader, key, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
	}

	s := hex.EncodeToString(signature)

	// try to VerifyPKCS1v15
	ss, err := hex.DecodeString(s)
	err = rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, hashed[:], ss)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
	}

	return s
}

func (card *CardBlock) build() {
	card.Version = version
	card.Hard = strconv.Itoa(hard)
	card.Timestamp = timestamp()
	card.RandNumber = randNumber()
	// 使用用户公钥，时间戳以及随机数作为种子
	key := card.PubKey + card.Timestamp + card.RandNumber
	// 去生成一个hash值，这里使用sha256这个比较公允的算法
	rawOre := sha256.Sum256([]byte(key))
	// 根据规则去判断hash是否是一张卡
	card.CardID = findCard(rawOre)
}

func (card *CardBlock) json() string {
	json, _ := json.Marshal(card)
	return string(json)
}

func (card CardBlock) cut(from int, to int) string {
	return string(card.CardID[len(card.CardID)+from : len(card.CardID)+to])
}

func (card CardBlock) cid() (int, error) {
	raw := card.cut(-7, -4)
	p1, error := strconv.Atoi(raw[0:1])
	p2, error := strconv.Atoi(raw[1:2])
	p3, error := strconv.Atoi(raw[2:3])
	return p1 + p2 + p3, error
}

func (card CardBlock) card() (Card, error) {
	id, error := card.cid()
	attack, error := strconv.Atoi(card.cut(-4, -2))
	defense, error := strconv.Atoi(card.cut(-2, 0))
	c := Card{
		id:      id,
		attack:  attack,
		defense: defense,
	}
	return c, error
}
