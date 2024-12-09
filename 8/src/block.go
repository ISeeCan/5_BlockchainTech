package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "strconv"
    "strings"
    "time"
)

// 定义区块结构
type Block struct {
    Lasthash  string
    Hash      string
    Data      string
    Timestamp string
    Height    int
    DiffNum   uint
    Nonce     int64
}

func (b *Block) getHash() string {
    hashData := fmt.Sprintf("%s%s%s%d%d%d", b.Lasthash, b.Data, b.Timestamp, b.Height, b.DiffNum, b.Nonce)
    hash := sha256.Sum256([]byte(hashData))
    return hex.EncodeToString(hash[:])
}

// 定义区块链
var blockchain []Block

func mine(data string) Block {
    if len(blockchain) < 1 {
        log.Panic("还未生成创世区块！")
    }

    lastBlock := blockchain[len(blockchain)-1]
    newBlock := Block{
        Lasthash:  lastBlock.Hash,
        Data:      data,
        Timestamp: time.Now().String(),
        Height:    lastBlock.Height + 1,
        DiffNum:   2, // 假设难度值为 2
    }

    for {
        newBlock.Nonce++
        newBlock.Hash = newBlock.getHash()
        // 将 newBlock.DiffNum 转换为 int
        if strings.HasPrefix(newBlock.Hash, strings.Repeat("0", int(newBlock.DiffNum))) {
            break
        }
    }

    return newBlock
}

func main() {
    genesisBlock := &Block{
        Timestamp: time.Now().String(),
        Data:      "我是创世区块！",
        Lasthash:  "0000000000000000000000000000000000000000000000000000000000000000",
        Height:    1,
        Nonce:     0,
        DiffNum:   2,
    }
    genesisBlock.Hash = genesisBlock.getHash()
    fmt.Println(*genesisBlock)

    blockchain = append(blockchain, *genesisBlock)

    for i := 0; i < 10; i++ {
        newBlock := mine("天气不错" + strconv.Itoa(i))
        blockchain = append(blockchain, newBlock)
        fmt.Println(newBlock)
    }
}
