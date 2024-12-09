// ADD all funcs
package main

import (
	"fmt"
    "strconv"
)

type Blockchain struct {
    blocks []*Block
}


func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks)-1]
    newBlock := NewBlock(data, prevBlock.Hash)
    bc.blocks = append(bc.blocks, newBlock)
}


func NewGenesisBlock() *Block {
    return NewBlock("Genesis Block", []byte{})
}


func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}

//4-2-2
// func main() {
//     bc := NewBlockchain()

//     bc.AddBlock("Send 1 BTC to Ivan")
//     bc.AddBlock("Send 2 more BTC to Ivan")

//     for _, block := range bc.blocks {
//         fmt.Printf("PrevHash: %x\n", block.PrevHash)
//         fmt.Printf("Data: %s\n", block.Data)
//         fmt.Printf("Hash: %x\n", block.Hash)
//         fmt.Println()
//     }
// }

//4-2-3
func main() {
    bc := NewBlockchain()

    bc.AddBlock("Send 1 BTC to Ivan")
    bc.AddBlock("Send 2 more BTC to Ivan")

    for _, block := range bc.blocks {
        fmt.Printf("PrevHash: %x\n", block.PrevHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.Hash)
        pow := NewProofOfWork(block)
        fmt.Printf("PoW Valid: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }
}
