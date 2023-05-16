package simple_block_chain

import (
	"fmt"
	"time"
)

type Blockchain []Block

func NewBlockchain() *Blockchain {
	genesisBlock := Block{}
	genesisBlock = Block{0, time.Now().String(), "", calculateHash(genesisBlock), "", difficulty}
	bc := make(Blockchain, 0)
	bc = append(bc, genesisBlock)
	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	previousBlock := (*bc)[len(*bc)-1]
	newBlock := createBlock(previousBlock, data)
	*bc = append(*bc, newBlock)
}

func createBlock(previousBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = previousBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	newBlock = mine(newBlock)
	return newBlock
}

func (bc *Blockchain) ValidateChain() bool {
	for i := 1; i < len(*bc); i++ {
		if (*bc)[i].PrevHash != (*bc)[i-1].Hash {
			return false
		}
		if calculateHash((*bc)[i]) != (*bc)[i].Hash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) ReplaceChain(newBC *Blockchain) {
	if len(*newBC) > len(*bc) && newBC.ValidateChain() {
		*bc = *newBC
	} else {
		fmt.Println("Received blockchain is not valid.")
	}
}
