package main

import (
	"encryptcard/protoc/cardproto"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

// StartClient 启动客户端
func StartClient() {
	queue := cellnet.NewEventQueue()
	peer := socket.NewConnector(queue)
	peer.Start("127.0.0.1:22366")
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

// RequestCardBlocksFetch 请求服务从height获得区块
func RequestCardBlocksFetch(peer cellnet.Peer, height int64) {
	CurrentSession(peer).Send(&cardproto.CardBlocksFetchRequest{
		Height: height,
	})
}
