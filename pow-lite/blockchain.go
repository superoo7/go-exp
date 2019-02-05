package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// GenerateGenesisBlock generate genesis block
func GenerateGenesisBlock() {
	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), "", GenerateHash(genesisBlock), "", Difficulty, ""}

	lock.Lock()
	Blockchain = append(Blockchain, genesisBlock)
	lock.Unlock()
}

// GenerateHash generate sha256 hash of a block
func GenerateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Msg + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock Generate a block into blockchain
func GenerateBlock(oldBlock Block, msg string) Block {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Msg = msg
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = Difficulty
	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHashValid(GenerateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(GenerateHash(newBlock), " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(GenerateHash(newBlock), " work done!")
			newBlock.Hash = GenerateHash(newBlock)
			break
		}
	}
	return newBlock
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}

// isBlockValid make sure block is valid by checking index, and comparing the hash of the previous block
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if GenerateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
