package blockchain

import (
	"fmt"
	"strings"
)

type BlockChain struct {
	TransactionPool []*Transaction
	Chain           []*Block
	BlockChainAddr  string
}

func NewBlockChain(blockAddr string) *BlockChain {
	b := &Block{}
	bc := &BlockChain{}
	bc.BlockChainAddr = blockAddr
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *BlockChain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	//creating new block
	b := NewBlock(nonce, prevHash, bc.TransactionPool)

	//adding the new block to the chain
	bc.Chain = append(bc.Chain, b)
	//clearing the transaction Pool
	bc.TransactionPool = []*Transaction{}

	return b
}

// getting the last block on the chain
func (bc *BlockChain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

// printing the block chain
func (bc *BlockChain) Print() {
	for i, b := range bc.Chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		b.Print()
		fmt.Printf("%s\n", strings.Repeat("*", 25))
	}
}
