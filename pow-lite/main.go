package main

import (
	"fmt"
	"log"
	"sync"
)

// ===CONFIG===
// lock
var lock sync.Mutex

// Difficulty Mining difficulty
const Difficulty = 1

// Port for tcp
const Port = 8081

// Blockchain array of blocks
var Blockchain []Block

func main() {
	go GenerateGenesisBlock()
	log.Fatal(run())
}

func run() error {
	fmt.Printf("Server running at http://localhost:%d\n", Port)
	if err := RunServer(); err != nil {
		return err
	}
	return nil
}
