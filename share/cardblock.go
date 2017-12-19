package share

import (
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
const version = "v0.2.0"

// CardBlock is a good block
type CardBlock struct {
	Version    string
	Hard       string
	PubKey     string
	Timestamp  string
	RandNumber string
	PrevCardID string
}

// Card is CardBlock helper
type Card struct {
	ID      int
	Attack  int
	Defense int
	CardID  string
}

// 时间戳
func timestamp() string {
	timestamp := time.Now().UnixNano()
	return strconv.Itoa(int(timestamp))
}

// 随机数
func randNumber() string {
	return strconv.Itoa(rand.Intn(1000))
}

// 判定是否为卡的函数
func findCard(card string, hard int) string {
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
	if err != nil || (i == 0 && last[0:2] != "000") {
		return ""
	}

	return card
}

// func (card *CardBlock) sign(key *rsa.PrivateKey) string {
// 	m := card.Version + card.PubKey + card.Timestamp + card.RandNumber +
// 		card.Hard + card.CardID
// 	message := []byte(m)

// 	// Only small messages can be signed directly; thus the hash of a
// 	// message, rather than the message itself, is signed. This requires
// 	// that the hash function be collision resistant. SHA-256 is the
// 	// least-strong hash function that should be used for this at the time
// 	// of writing (2016).
// 	hashed := sha256.Sum256(message)

// 	signature, err := rsa.SignPKCS1v15(crand.Reader, key, crypto.SHA256, hashed[:])
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
// 	}

// 	s := hex.EncodeToString(signature)

// 	// try to VerifyPKCS1v15
// 	ss, err := hex.DecodeString(s)
// 	err = rsa.VerifyPKCS1v15(&key.PublicKey, crypto.SHA256, hashed[:], ss)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
// 	}

// 	return s
// }

// VerifyCardID 验证卡片id，也就是验证区块是不是真的
func (card *CardBlock) VerifyCardID() bool {
	if findCard(card.CardID(), card.HardInt()) == "" {
		fmt.Fprintf(os.Stderr, "CardID验证失败，你可能是一张假卡")
		return false
	}
	return true
}

// 验证签名
// func (card *CardBlock) verifySign() bool {
// 	m := card.Version + card.PubKey + card.Timestamp + card.RandNumber +
// 		card.Hard + card.CardID
// 	message := []byte(m)

// 	hashed := sha256.Sum256(message)

// 	// try to VerifyPKCS1v15
// 	ss, err := hex.DecodeString(card.Signature)

// 	key, err := loadPublicKeyFromString(card.PubKey)

// 	err = rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed[:], ss)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
// 		return false
// 	}

// 	return true
// }

// Verify 验证是否为真卡
func (card *CardBlock) Verify() bool {
	return card.VerifyCardID()
}

// LoadCard json to CardBlock
func LoadCard(filePath string) CardBlock {
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

// HardInt Hard to int
func (card *CardBlock) HardInt() int {
	i, err := strconv.ParseInt(card.Hard, 10, 32)
	if err != nil {
		return 0
	}
	return int(i)
}

// CardID 获取区块的ID
func (card *CardBlock) CardID() string {
	// 使用用户公钥，时间戳以及随机数作为种子
	key := card.Version + card.PubKey + card.Timestamp +
		card.RandNumber + card.Hard + card.PrevCardID
	// 去生成一个hash值，这里使用sha256这个比较公允的算法
	hashCard := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hashCard[:])
}

// Build 构建一张卡
func (card *CardBlock) Build() string {
	card.Version = version
	card.Timestamp = timestamp()
	card.RandNumber = randNumber()
	// 根据规则去判断hash是否是一张卡
	return findCard(card.CardID(), card.HardInt())
}

// JSON 区块json
func (card *CardBlock) JSON() string {
	json, _ := json.Marshal(card)
	return string(json)
}

// Cut 切区块字符串
func (card CardBlock) Cut(from int, to int) string {
	CardID := card.CardID()
	return string(CardID[len(CardID)+from : len(CardID)+to])
}

// Cid 分析区块的卡片id
func (card CardBlock) Cid() (int, error) {
	raw := card.Cut(-7, -4)
	p1, error := strconv.Atoi(raw[0:1])
	p2, error := strconv.Atoi(raw[1:2])
	p3, error := strconv.Atoi(raw[2:3])
	return p1 + p2 + p3, error
}

// Card 分析区块包含的卡
func (card CardBlock) Card() (Card, error) {
	fmt.Printf("CardID: %s\n", card.CardID())
	id, error := card.Cid()
	attack, error := strconv.Atoi(card.Cut(-4, -2))
	defense, error := strconv.Atoi(card.Cut(-2, 0))
	c := Card{
		ID:      id,
		Attack:  attack,
		Defense: defense,
	}
	return c, error
}
