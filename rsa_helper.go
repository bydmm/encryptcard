package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const privatePath = "./saves/private.pem"
const publicPath = "./saves/public.pem"

// 从磁盘获取密钥，如果得不到就重新生成
func getKeyPair() *rsa.PrivateKey {
	key, err := getKeyPairFromDisk()
	if err != nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("找不到密钥对, 输入 [Y] 创建密钥对 [N]: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		if strings.ToUpper(text) == "Y" {
			return generateRSAKeys()
		}
		checkError(err)
	}
	return key
}

// 从磁盘读公钥
func loadPublicKeyFromDisk() (*rsa.PublicKey, error) {
	raw, err := ioutil.ReadFile(publicPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(raw)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		// fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
		return pub, nil
	default:
		panic("unknown type of public key")
	}
}

// 从string读公钥
func loadPublicKeyFromString(key string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		// fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
		return pub, nil
	default:
		panic("unknown type of public key")
	}
}

// 我今天才知道，私钥可以算公钥，有私钥删掉公钥都行
func loadPrivateKeyFromDisk() (*rsa.PrivateKey, error) {
	raw, err := ioutil.ReadFile(privatePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(raw)

	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil || block.Type != "PRIVATE KEY" {
		// fmt.Printf("Got a %T, with remaining data: %q", key, rest)
		return key, nil
	}

	return nil, fmt.Errorf("Failed to parse private key")
}

// 从磁盘获得钥匙的对象。实际上只读私钥就行了
func getKeyPairFromDisk() (*rsa.PrivateKey, error) {
	pubKey, err := loadPublicKeyFromDisk()
	if pubKey == nil && err != nil {
		return nil, fmt.Errorf("Failed to find %s", publicPath)
	}
	privateKey, err := loadPrivateKeyFromDisk()
	if err != nil {
		return nil, fmt.Errorf("Failed to find %s", privatePath)
	}
	// privateKey.PublicKey = *pubKey
	return privateKey, nil
}

// 生成RSA钥匙对
func generateRSAKeys() *rsa.PrivateKey {
	reader := rand.Reader
	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	savePEMKey(privatePath, key)
	savePublicPEMKey(publicPath, key.PublicKey)
	checkError(err)
	return key
}

// 返回公钥的string，要填写在区块用的
func pubKey(key *rsa.PrivateKey) string {
	PubASN1, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	checkError(err)

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})

	return string(pubBytes)
}

// 把公钥存盘
func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	PubASN1, err := x509.MarshalPKIXPublicKey(&pubkey)
	checkError(err)

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})

	ioutil.WriteFile(fileName, pubBytes, 0644)
}

// 把私钥存盘
func savePEMKey(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
