package share

import (
	"errors"
	"time"
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

// AdaptiveHard 根据之前区块被挖掘的速度去调整新块产生的速度
// 一分钟出块，如果低于一分钟就加难度，高于两分钟就减难度
// 无上限，无减半
func (chain *CardBlockChain) AdaptiveHard() int32 {
	if chain.Height() < 2 {
		return 0
	}
	head := chain.HeadBlock()
	secondHead, _ := chain.BlockAtHeight(head.Height - 1)
	blockTime := time.Duration(head.Timestamp - secondHead.Timestamp)
	switch {
	case blockTime < (1 * time.Minute):
		return head.Hard + 1
	case blockTime > (2 * time.Minute):
		return head.Hard + 1
	default:
		return head.Hard
	}

}
