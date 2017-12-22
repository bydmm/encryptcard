package main

import (
	"encryptcard/protoc/cardproto"
	"encryptcard/share"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
)

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
	})
}

// RequestCardBlocksFetch 请求服务器从height发送区块
func RequestCardBlocksFetch(peer cellnet.Peer, height int64) {
	CurrentSession(peer).Send(&cardproto.CardBlocksFetchRequest{
		Height: height,
	})
}

// CardBlockFetchResponse 从服务器获得连续区块
func CardBlockFetchResponse(peer cellnet.Peer, userCallback func(*cardproto.CardBlockFetchResponse)) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlockFetchResponse", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockFetchResponse)
		userCallback(msg)
	})

}

// CardBlockLiveMsg 服务器推送区块
func CardBlockLiveMsg(peer cellnet.Peer, userCallback func(*cardproto.CardBlock)) {
	cellnet.RegisterMessage(peer, "cardproto.CardBlock", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlock)
		userCallback(msg)
	})
}

// RequestCardBlockPush 把挖到的区块发给服务器
func RequestCardBlockPush(peer cellnet.Peer, cardBlock *share.CardBlock) {
	CurrentSession(peer).Send(&cardproto.CardBlockPushRequest{
		CardBlock: share.CardBlockToMsg(cardBlock),
	})
}

// StartClient 启动客户端
func StartClient(userCallback func(queue cellnet.EventQueue, peer cellnet.Peer, ev *cellnet.Event, success bool)) {
	queue := cellnet.NewEventQueue()
	peer := socket.NewConnector(queue)
	peer.Start("47.52.235.110:22366")

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
