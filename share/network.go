package share

import (
	"encryptcard/protoc/cardproto"
)

// CardBlockToMsg 从区块到消息
func CardBlockToMsg(block *CardBlock) *cardproto.CardBlock {
	return &cardproto.CardBlock{
		Version:    block.Version,
		Hard:       block.Hard,
		PubKey:     block.PubKey,
		Timestamp:  block.Timestamp,
		RandNumber: block.RandNumber,
		PrevCardID: block.PrevCardID,
		Height:     block.Height,
	}
}

// ProtoToBlock 从消息到区块
func ProtoToBlock(msg *cardproto.CardBlock) *CardBlock {
	return &CardBlock{
		Version:    msg.Version,
		Hard:       msg.Hard,
		PubKey:     msg.PubKey,
		Timestamp:  msg.Timestamp,
		RandNumber: msg.RandNumber,
		PrevCardID: msg.PrevCardID,
		Height:     msg.Height,
	}
}

// PushProtoToBlock 解析提交区块消息
func PushProtoToBlock(msg *cardproto.CardBlockPushRequest) *CardBlock {
	return ProtoToBlock(msg.CardBlock)
}
