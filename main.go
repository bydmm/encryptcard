package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

func generateRSAKeys() *rsa.PrivateKey {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)
	return key
}

func pubKey(key *rsa.PrivateKey) string {
	asn1Bytes, err := asn1.Marshal(key.PublicKey)
	checkError(err)
	data := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	return string(pem.EncodeToMemory(data))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func main() {
	// 用户公钥

	key := generateRSAKeys()

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
