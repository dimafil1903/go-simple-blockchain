package simple_block_chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

const difficulty = 1

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func mine(block Block) Block {
	for i := 0; ; i++ {
		block.Nonce = i
		block.Hash = calculateHash(block)
		if isValidHash(block.Hash) {
			break
		}
	}
	return block
}
func isValidHash(hash string) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix)
}
