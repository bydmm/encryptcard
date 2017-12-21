package share

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)

// Store 存bin
func Store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// Load 读bin
func Load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

// LoadCardBlockChainFromDisk 从文件读取历史区块
func LoadCardBlockChainFromDisk(path string) *CardBlockChain {
	var cardChainRead CardBlockChain
	Load(&cardChainRead, path)
	return &cardChainRead
}
