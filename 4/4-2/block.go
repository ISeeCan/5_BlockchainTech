// ADD sethash
package main

import (
	//"crypto/sha256"
	//"encoding/binary"
	//"bytes"
	"time"
)

// before 4-2-3
// type Block struct {
// 	Time     int64
// 	PrevHash []byte
// 	Hash     []byte
// 	Data     []byte
// }

//4-2-3
type Block struct {
	Time     int64
	PrevHash []byte
	Hash     []byte
	Data     []byte
	Nonce int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Time:     time.Now().Unix(),
		PrevHash: prevHash,
		Data:     []byte(data),
		Hash:     []byte{},
	}
	block.SetHash()
	return block
}

// before 4-2-3
// func (b *Block) SetHash() {
// 	// 将 Time 转换为字节数组
// 	timeBytes := make([]byte, 8)
// 	binary.BigEndian.PutUint64(timeBytes, uint64(b.Time))

// 	// 拼接 PrevHash、Time、Data
// 	headers := bytes.Join([][]byte{b.PrevHash, timeBytes, b.Data}, []byte{})

// 	// 计算 SHA-256 哈希
// 	hash := sha256.Sum256(headers)
// 	b.Hash = hash[:]
// }

//4-2-3
func (b *Block) SetHash() {
    pow := NewProofOfWork(b)
    nonce, hash := pow.Run()
    b.Hash = hash[:]
    b.Nonce = nonce
}