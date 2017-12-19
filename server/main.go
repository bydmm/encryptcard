package main

import (
	"encryptcard/protoc/cardproto"
	"encryptcard/share"
	"fmt"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/socket"
	"github.com/davyxu/golog"
)

// CreationBLOCKID 创始区块
const CreationBLOCKID = "de84191c7de8ca2ae1cb9864046b453f42100cda096ef1fd12e72edea0009999"

var log = golog.New("main")

func blockToMessage(block share.CardBlock) cardproto.CardBlockReceive {
	return cardproto.CardBlockReceive{
		Version:    block.Version,
		Hard:       block.Hard,
		PubKey:     block.PubKey,
		Timestamp:  block.Timestamp,
		RandNumber: block.RandNumber,
		PrevCardID: block.PrevCardID,
	}
}

func main() {
	queue := cellnet.NewEventQueue()

	peer := socket.NewAcceptor(queue).Start("127.0.0.1:22366")
	peer.SetName("client")

	path := fmt.Sprintf("./blocks/%s.json", CreationBLOCKID)
	headBlock := share.LoadCard(path)

	fmt.Printf("CardID: %s", headBlock.CardID())

	cellnet.RegisterMessage(peer, "cardproto.CardBlockFetch", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockFetch)
		fmt.Printf("CardBlockFetch: %s", msg.CardID)
		ack := blockToMessage(headBlock)
		ev.Ses.Send(&ack)
	})

	cellnet.RegisterMessage(peer, "cardproto.CardBlockDig", func(ev *cellnet.Event) {
		msg := ev.Msg.(*cardproto.CardBlockDig)

		block := share.CardBlock{
			Version:    msg.Version,
			Hard:       msg.Hard,
			PubKey:     msg.PubKey,
			Timestamp:  msg.Timestamp,
			RandNumber: msg.RandNumber,
			PrevCardID: msg.PrevCardID,
		}

		if block.VerifyCardID() {
			headBlock = block
			ack := blockToMessage(headBlock)

			// 广播给所有连接
			peer.VisitSession(func(ses cellnet.Session) bool {

				ses.Send(&ack)

				return true
			})
		}
	})

	queue.StartLoop()

	queue.Wait()

	peer.Stop()
}
