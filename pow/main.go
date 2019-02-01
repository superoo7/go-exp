package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// Block store data into a block
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// BlockChain core struct for bc
type BlockChain struct {
	blocks []*Block
}

// DeriveHash methods
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock create a block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// addBlock add block to blockchain
func (chain *BlockChain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

// Genesis create genesis block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockchain initialize blockchain with genesis block
func InitBlockchain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockchain()

	chain.addBlock("something 1")
	chain.addBlock("something 2")
	chain.addBlock("something 3")
	chain.addBlock("something 4")

	for _, block := range chain.blocks {
		fmt.Printf("PrevHash : %x\n", block.PrevHash)
		fmt.Printf("Data     : %s\n", block.Data)
		fmt.Printf("Hash     : %x\n", block.Hash)
		fmt.Println("===========")
	}
}
