//Changed Nothing
package main

import (
	"fmt"
	"time"
)

// 4-2-1
func main() {
	t := time.Now()

	block := NewBlock("Genesis Block", []byte{})

	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	fmt.Printf("Time: %s\n", time.Unix(block.Time, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)

	fmt.Println("Time using: ", time.Since(t))
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
