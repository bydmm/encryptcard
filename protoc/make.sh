#!/bin/sh

go build -o ./protoc-gen-msg github.com/davyxu/cellnet/protoc-gen-msg
go build -o ./protoc-gen-go github.com/golang/protobuf/protoc-gen-go

protoc --plugin=protoc-gen-go=protoc-gen-go --go_out=./cardproto  *.proto
protoc --plugin=protoc-gen-msg=protoc-gen-msg --msg_out=msgid.go:./cardproto *.proto