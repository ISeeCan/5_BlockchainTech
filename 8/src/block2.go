package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// 区块结构体
type Block struct {
	Prehash   string
	Hash      string
	Timestamp string
	Data      string
	Height    int
	Address   string
}

// 挖矿节点结构体
type Node struct {
	Tokens  int
	Days    int
	Address string
}

// 全局变量
var blockchain []Block                // 区块链
var mineNodesPool []Node              // 挖矿节点池
var probabilityNodesPool []string     // 随机节点池，用地址填充

// 初始化函数
func init() {
	// 手动添加两个节点
	mineNodesPool = append(mineNodesPool, Node{1000, 1, "AAAAAAAAAA"})
	mineNodesPool = append(mineNodesPool, Node{100, 3, "BBBBBBBBBB"})

	// 根据节点权重初始化随机节点池
	for _, v := range mineNodesPool {
		for i := 0; i < v.Tokens*v.Days; i++ {
			probabilityNodesPool = append(probabilityNodesPool, v.Address)
		}
	}
	fmt.Println("随机节点池长度:", len(probabilityNodesPool))
}

// 生成区块哈希
func (b *Block) getHash() {
	record := b.Prehash + b.Timestamp + b.Data + b.Address
	hash := sha256.Sum256([]byte(record))
	b.Hash = hex.EncodeToString(hash[:])
}

// 获取随机节点地址
func getMineNodeAddress() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(probabilityNodesPool))
	return probabilityNodesPool[index]
}

// 生成新区块
func generateNewBlock(previousBlock Block, data string) Block {
	newBlock := Block{
		Prehash:   previousBlock.Hash,
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      data,
		Height:    previousBlock.Height + 1,
		Address:   getMineNodeAddress(),
	}
	newBlock.getHash()
	return newBlock
}

// 主函数
func main() {
	// 初始化系统 自动调用，无需手动
	//init()

	// 创建创世区块
	genesisBlock := Block{
		Prehash:   "0000000000000000000000000000000000000000000000000000000000000000",
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Data:      "我是创世区块",
		Height:    1,
		Address:   "0000000000",
	}
	genesisBlock.getHash()

	// 将创世区块添加到区块链中
	blockchain = append(blockchain, genesisBlock)
	fmt.Println("创世区块:", genesisBlock)

	// 挖掘新区块
	i := 0
	for {
		time.Sleep(1 * time.Second) // 模拟出块间隔
		newBlock := generateNewBlock(blockchain[i], fmt.Sprintf("我是第 %d 个区块", i+1))
		blockchain = append(blockchain, newBlock)
		fmt.Println("新区块:", newBlock)
		i++
	}
}
