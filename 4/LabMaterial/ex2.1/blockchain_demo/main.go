package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	block := NewBlock("Genesis Block", []byte{})

	fmt.Printf("Prev. hash: %x\n", block.PrevHash)
	fmt.Printf("Time: %s\n", time.Unix(block.Time, 0).Format("2006-01-02 15:04:05"))
	fmt.Printf("Data: %s\n", block.Data)
	fmt.Printf("Hash: %x\n", block.Hash)

	fmt.Println("Time using: ", time.Since(t))
}
