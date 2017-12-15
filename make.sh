#!/bin/sh

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o encryptcard_linux_x64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o encryptcard_win_x64.exe
go build -o  encryptcard_darwin_x64
