package main

import (
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// 难度系数
const hard int = 3
const version = "v0.0.1"

// CardBlock is a good block
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

// 验证卡片id，也就是验证区块是不是真的
func (card *CardBlock) verifyCardID() bool {
	key := card.Version + card.PubKey + card.Timestamp + card.RandNumber + card.Hard

	rawOre := sha256.Sum256([]byte(key))

	if card.CardID != findCard(rawOre) {
		fmt.Fprintf(os.Stderr, "CardID验证失败，你可能是一张假卡")
		return false
	}
	return true
}

// 验证签名
func (card *CardBlock) verifySign() bool {
	m := card.Version + card.PubKey + card.Timestamp + card.RandNumber +
		card.Hard + card.CardID
	message := []byte(m)

	hashed := sha256.Sum256(message)

	// try to VerifyPKCS1v15
	ss, err := hex.DecodeString(card.Signature)

	key, err := loadPublicKeyFromString(card.PubKey)

	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed[:], ss)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return false
	}

	return true
}

// 验证是否为真卡
func (card *CardBlock) verify() bool {
	return card.verifySign() && card.verifyCardID()
}

// 去读卡片json
func loadCard(filePath string) CardBlock {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var card CardBlock
	if err := json.Unmarshal(raw, &card); err != nil {
		panic(err)
	}

	return card
}

func (card *CardBlock) build() {
	card.Version = version
	card.Hard = strconv.Itoa(hard)
	card.Timestamp = timestamp()
	card.RandNumber = randNumber()
	// 使用用户公钥，时间戳以及随机数作为种子
	key := card.Version + card.PubKey + card.Timestamp + card.RandNumber + card.Hard
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
