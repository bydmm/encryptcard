package main

import (
	"encryptcard/share"
	"fmt"
)

// CreationBLOCKID 创始区块ID
const CreationBLOCKID = "0ee56b47756d5d7c04aa0270b601a04cbf25a81a06e2775432e32b83e0009999"

// ChainPath 区块链文件路径
const ChainPath = "./cards.chain"

// CreationBLOCK 读取创始区块
func CreationBLOCK() *share.CardBlock {
	path := fmt.Sprintf("./blocks/%s.json", CreationBLOCKID)
	block := share.LoadCard(path)
	return &block
}

// NewCardBlockChain 新建一条主链
func NewCardBlockChain() *share.CardBlockChain {
	var blocks []*share.CardBlock
	blocks = append(blocks, CreationBLOCK())
	return &share.CardBlockChain{Cardblocks: blocks}
}

// 初始化区块链
func main() {
	cardChain := NewCardBlockChain()
	share.Store(cardChain, ChainPath)

	var cardChainRead share.CardBlockChain
	share.Load(&cardChainRead, ChainPath)

	creationBLOCK := cardChainRead.Cardblocks[0]
	fmt.Printf("CardID: %s\n", creationBLOCK.CardID())
	fmt.Printf("%s\n", creationBLOCK.PubKey)
}
