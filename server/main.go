package main

import (
	"encryptcard/protoc/cardproto"
	"encryptcard/share"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
	"github.com/davyxu/golog"
)

// CreationBLOCKID 创始区块
const CreationBLOCKID = "0ee56b47756d5d7c04aa0270b601a04cbf25a81a06e2775432e32b83e0009999"

// ChainPath 区块链文件路径
const ChainPath = "./blocks/cards.chain"

// OnceSyncHeight 同步区块高度一次读取多少个
const OnceSyncHeight = 20

// 初始化logging
var log = golog.New("main")

// cardblockChain 内存中的区块链
var cardblockChain *share.CardBlockChain

// BlockToMessage 从区块到消息
func BlockToMessage(block *share.CardBlock) *cardproto.CardBlock {
	return &cardproto.CardBlock{
		Version:    block.Version,
		Hard:       block.Hard,
		PubKey:     block.PubKey,
		Timestamp:  block.Timestamp,
		RandNumber: block.RandNumber,
		PrevCardID: block.PrevCardID,
	}
}

// MessageToBlock 从消息到区块
func MessageToBlock(msg cardproto.CardBlockPushRequest) *share.CardBlock {
	return &share.CardBlock{
		Version:    msg.CardBlock.Version,
		Hard:       msg.CardBlock.Hard,
		PubKey:     msg.CardBlock.PubKey,
		Timestamp:  msg.CardBlock.Timestamp,
		RandNumber: msg.CardBlock.RandNumber,
		PrevCardID: msg.CardBlock.PrevCardID,
	}
}

// LoadCardBlockChainFromDisk 从文件读取历史区块
func LoadCardBlockChainFromDisk() *share.CardBlockChain {
	var cardChainRead *share.CardBlockChain
	share.Load(cardChainRead, ChainPath)
	return cardChainRead
}

// AddBlockToChain 向链条加块
func AddBlockToChain(block *share.CardBlock) {
	cardblockChain.Cardblocks = append(cardblockChain.Cardblocks, block)
}

// 初始化区块链
func initChain() {
	log.Infof("从磁盘读取区块链....\n")
	cardblockChain = LoadCardBlockChainFromDisk()
	log.Infof("区块链读取完成....\n")
	log.Infof("区块高度: %d\n", len(cardblockChain.Cardblocks))
	log.Infof("区块高度: %d\n", len(cardblockChain.Cardblocks))
}

// 初始化
func initServer() {
	// 初始化区块链
	initChain()
}

// CardBlockSyncRequest 用户同步请求
func CardBlockSyncRequest(peer cellnet.Peer) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlockSyncRequest", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockSyncRequest)

		// 判断用户是否跟随了主链
		valid := false
		// 先判断高度是否吻合
		if msg.Height <= cardblockChain.Height() {
			cardBlock := cardblockChain.Cardblocks[msg.Height]
			if cardBlock != nil {
				// 判断区块ID是否一致
				valid = (msg.CardID == cardBlock.CardID())
			} else {
				panic("区块链损坏??")
			}
		}

		res := cardproto.CardBlockSyncResponse{
			Valid:      valid,
			Height:     cardblockChain.Height(),
			CardID:     cardblockChain.HeadBlock().CardID(),
			PrevCardID: cardblockChain.HeadBlock().PrevCardID,
		}

		ev.Ses.Send(&res)
	})
}

// CardBlocksFetchRequest 用户获取区块请求
func CardBlocksFetchRequest(peer cellnet.Peer) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlocksFetchRequest", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlocksFetchRequest)

		valid := false
		var msgBlocks []*cardproto.CardBlock

		// 判断区块高度是否正确
		if msg.Height <= cardblockChain.Height() {
			valid = true
			index := 0
			CurrentHeight := msg.Height
			for index < OnceSyncHeight {
				CurrentHeight++
				msgBlock, err := cardblockChain.BlockAtHeight(CurrentHeight)
				// 判断如果该块不存在，跳过
				if err == nil {
					continue
				}
				msgBlocks = append(msgBlocks, BlockToMessage(msgBlock))
				index++
			}
		}

		res := cardproto.CardBlockFetchResponse{
			Valid:      valid,
			CardBlocks: msgBlocks,
		}

		ev.Ses.Send(&res)

	})
}

// CardBlockPushRequest 用户竞争上传区块
func CardBlockPushRequest(peer cellnet.Peer) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlockPushRequest", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockPushRequest)
		block := MessageToBlock(*msg)
		headBlock := cardblockChain.HeadBlock()

		// 不是一条链，拜拜
		if block.PrevCardID != headBlock.CardID() {
			return
		}

		if block.VerifyCardID() {
			//把新块加到链上
			AddBlockToChain(block)

			// 广播给所有连接, 新快已经产生了
			peer.VisitSession(func(ses cellnet.Session) bool {
				ses.Send(BlockToMessage(block))
				return true
			})
		}
	})
}

// 启动服务
func startServer() {
	queue := cellnet.NewEventQueue()
	peer := socket.NewAcceptor(queue).Start("0.0.0.0:22366")
	peer.SetName("client")

	// 用户同步区块的请求
	CardBlockSyncRequest(peer)

	// 用户读取区块的请求
	CardBlocksFetchRequest(peer)

	// 接受抢占式的块提交（挖矿）
	CardBlockPushRequest(peer)

	queue.StartLoop()
	queue.Wait()
}

func main() {
	initServer()
	startServer()
}
