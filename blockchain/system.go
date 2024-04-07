package blockchain

import (
	"fmt"
	"strings"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

// making a guessing function for the nodes to validate their work
func (bc *BlockChain) ValidProof(nonce int, prevHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{Timestamp: 0, Nonce: nonce, PrevHash: prevHash, Transactions: transactions}
	guessHash := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHash[:difficulty] == zeros
}

// checking if the validation is correct and returning the the challenge (nonce)
func (bc *BlockChain) ProofOfWork() int {
	transactions := bc.CopyTransaction()
	prevHash := bc.LastBlock().PrevHash
	nonce := 0
	for !bc.ValidProof(nonce, prevHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

// a function that is used by the nodes to create a new block, ...and  reward the miner of that block
func (bc *BlockChain) Mining() bool {
	bc.AddTransactions(MINING_SENDER, bc.BlockChainAddr, MINING_REWARD)
	prevHash := bc.LastBlock().Hash()
	nonce := bc.ProofOfWork()
	bc.CreateBlock(nonce, prevHash)
	fmt.Println("action = mining , status = success")
	return true
}
