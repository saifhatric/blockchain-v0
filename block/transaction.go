package block

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/saifhatric/blockchain-v1/wallet"
)

type Transaction struct {
	Sender  string
	Reciver string
	Value   float32
}

// creating new transaction and adding it to the transaction Pool
func (bc *BlockChain) AddTransactions(wt *wallet.WTransaction, sender, reciver string, value float32) bool {
	t := &Transaction{Sender: sender, Reciver: reciver, Value: value}

	if sender == MINING_SENDER {
		bc.TransactionPool = append(bc.TransactionPool, t)
		return true
	}
	if Verify(wt) {
		bc.TransactionPool = append(bc.TransactionPool, t)
		return true
	} else {

		fmt.Println("ERROR: invalid transaction ! ")
	}
	return false
}

func Verify(wt *wallet.WTransaction) bool {
	s := wt.GenerateSignature()
	m, err := json.Marshal(wt)
	if err != nil {
		panic(err)
	}

	h := sha256.Sum256([]byte(m))
	valid := ecdsa.Verify(wt.SenderPublicKey, h[:], s.R, s.S)
	return valid
}

// coping the transactions, for the validation of proof_of_work
func (bc *BlockChain) CopyTransaction() []*Transaction {
	transactions := []*Transaction{}
	transactions = append(transactions, bc.TransactionPool...)
	return transactions
}

func (bc *BlockChain) CalculateAmount(blockchainAddr string) float32 {
	var total float32 = 0.0

	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			if blockchainAddr == t.Reciver {
				total += t.Value
			}
			if blockchainAddr == t.Sender {
				total -= t.Value
			}

		}
	}

	return total
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("_", 20))
	fmt.Printf("sender %s\n", t.Sender)
	fmt.Printf("reciver %s\n", t.Reciver)
	fmt.Printf("value %.2f\n", t.Value)
	fmt.Printf("%s\n", strings.Repeat("_", 20))
}
