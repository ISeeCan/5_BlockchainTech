package main

type Blockchain struct {
}

func (bc *Blockchain) AddBlock(data string) {
	//可能用到的函数：
	//	len(array)：获取数组长度
	//	append(array,b):将元素b添加至数组array末尾
}

func NewGenesisBlock() *Block {
	//创世区块前置哈希为空，Data为"Genesis Block"
}

func NewBlockchain() *Blockchain {
}
