package main

import (
	"time"
)

type Block struct {
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevHash, []byte{}}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	//为Block生成hash，使用sha256.Sum256(data []byte)函数
}
