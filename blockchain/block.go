package blockchain

import (
	"fmt"
	"time"
)

type Block struct {
	Timestamp    int64
	PrevHash     [32]byte
	Transactions []*Transaction
	Nonce        int
}

// initializing a new block
func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Timestamp:    time.Now().UnixNano(),
		PrevHash:     prevHash,
		Transactions: transactions,
		Nonce:        nonce,
	}
}

func (b *Block) Print() {
	fmt.Printf("timeStamp %d\n", b.Timestamp)
	fmt.Printf("prevHash %x\n", b.PrevHash)
	fmt.Printf("transactions %v\n", b.Transactions)
	fmt.Printf("nonce %d\n", b.Nonce)
	for _, t := range b.Transactions {
		t.Print()
	}

}
