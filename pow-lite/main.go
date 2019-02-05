package main

import (
	"log"
	"sync"
)

// ===DATAMODEL===

// Block storing block in a blockchain
type Block struct {
	Index      int
	Timestamp  string
	Msg        string
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

// Message basic data being stored in blockchain
type Message struct {
	Msg string
}

// ===CONFIG===
// lock
var lock sync.Mutex

// Difficulty Mining difficulty
const Difficulty = 1

// Port for tcp
const Port = ":8081"

// Blockchain array of blocks
var Blockchain []Block

func main() {
	go GenerateGenesisBlock()
	log.Fatal(run())
}

func run() error {
	RunServer()
	return nil
}
