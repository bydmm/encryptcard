package main

import (
	"encryptcard/protoc/cardproto"
	"encryptcard/share"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

// MessageToBlock 从消息到区块
func MessageToBlock(msg cardproto.CardBlock) *share.CardBlock {
	return &share.CardBlock{
		Version:    msg.Version,
		Hard:       msg.Hard,
		PubKey:     msg.PubKey,
		Timestamp:  msg.Timestamp,
		RandNumber: msg.RandNumber,
		PrevCardID: msg.PrevCardID,
	}
}

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

// CurrentSession 获取当前session
func CurrentSession(peer cellnet.Peer) cellnet.Session {
	return peer.(socket.Connector).DefaultSession()
}

// RequestCardBlockSync 请求同步到最新的区块
func RequestCardBlockSync(peer cellnet.Peer, height int64, CardID string) {
	CurrentSession(peer).Send(&cardproto.CardBlockSyncRequest{
		Height: height,
		CardID: CardID,
	})
}

// CardBlockSyncResponse 获得同步状态
func CardBlockSyncResponse(peer cellnet.Peer, userCallback func(*cardproto.CardBlockSyncResponse)) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlockSyncResponse", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockSyncResponse)
		userCallback(msg)
		// // 判断还在不在主链之上
		// if msg.Valid {
		// 	block := chain.HeadBlock()
		// 	// 如果当先就是最新块
		// 	if msg.CardID == block.CardID() {
		// 		// TO DO..... 大概要开始挖了
		// 		return
		// 	}
		// 	// TO DO..... 大概要去同步了
		// } else {
		// 	log.Fatalf("抱歉，你已与主链脱节，请删除链文件重新同步\n")
		// }
	})
}

// RequestCardBlocksFetch 请求服务器从height发送区块
func RequestCardBlocksFetch(peer cellnet.Peer, height int64) {
	CurrentSession(peer).Send(&cardproto.CardBlocksFetchRequest{
		Height: height,
	})
}

// CardBlockFetchResponse 从服务器获得区块
func CardBlockFetchResponse(peer cellnet.Peer, userCallback func(*cardproto.CardBlockFetchResponse)) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlockFetchResponse", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockFetchResponse)
		userCallback(msg)
	})
}

// RequestCardBlockPush 把挖到的区块发给服务器
func RequestCardBlockPush(peer cellnet.Peer, cardBlock *share.CardBlock) {
	CurrentSession(peer).Send(&cardproto.CardBlockPushRequest{
		CardBlock: BlockToMessage(cardBlock),
	})
}

// StartClient 启动客户端
func StartClient(userCallback func(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event, success bool)) {
	queue := cellnet.NewEventQueue()
	peer := socket.NewConnector(queue)
	peer.Start("127.0.0.1:22366")

	queue.StartLoop()

	// 当服务链接成功
	cellnet.RegisterMessage(peer, "coredef.SessionConnected", func(ev *cellnet.Event) {
		userCallback(queue, peer, ev, true)
	})

	// 如果网路断了
	cellnet.RegisterMessage(peer, "coredef.SessionConnectFailed", func(ev *cellnet.Event) {
		// 会话连接失败
		userCallback(queue, peer, ev, false)
	})

	// 设置连接超时2秒后自动重连
	peer.(socket.Connector).SetAutoReconnectSec(2)

	queue.Wait()
}
