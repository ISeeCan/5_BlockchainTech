// ADD some imports, ADD Run, Validate, Updated block.go->block
package main

import (
	"bytes"
	"math/big"

	"crypto/sha256"
    "fmt"
)

// It doesn't match the instruction...
const targetBits = 8

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToHex(pow.block.Time),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
		var hashInt big.Int
		var hash [32]byte
		nonce := 0
	
		fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
		
		for {
			// 准备数据
			data := pow.prepareData(nonce)
			// 计算哈希
			hash = sha256.Sum256(data)
			
			// 转换为big.Int类型以进行比较
			hashInt.SetBytes(hash[:])
	
			// 判断是否小于target
			if hashInt.Cmp(pow.target) == -1 {
				break
			} else {
				nonce++
			}
		}
		fmt.Printf("\nHash: %x\n", hash)
	
		return nonce, hash[:]
	
}

func (pow *ProofOfWork) Validate() bool {
		var hashInt big.Int
	
		// 准备数据并计算哈希
		data := pow.prepareData(pow.block.Nonce)
		hash := sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
	
		// 判断是否小于target
		isValid := hashInt.Cmp(pow.target) == -1
		return isValid
}
