package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"log"
)

// generating a hash for the block
func (b *Block) Hash() [32]byte {
	jBlock, err := json.Marshal(b)
	if err != nil {
		log.Fatal(err)
	}
	return sha256.Sum256(jBlock)
}
