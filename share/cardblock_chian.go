package share

import (
	"errors"
)

// CardBlockChain 区块链
type CardBlockChain struct {
	Cardblocks []*CardBlock
}

// HeadBlock 最后一个块
func (chain *CardBlockChain) HeadBlock() *CardBlock {
	index := len(chain.Cardblocks) - 1
	return chain.Cardblocks[index]
}

// Height 区块链高度
func (chain *CardBlockChain) Height() int64 {
	return int64(len(chain.Cardblocks) - 1)
}

// BlockAtHeight 获取特定高度的区块
func (chain *CardBlockChain) BlockAtHeight(height int64) (*CardBlock, error) {
	if height > chain.Height() {
		return nil, errors.New("Too height in Chain")
	}
	block := chain.Cardblocks[height]
	return block, nil
}
