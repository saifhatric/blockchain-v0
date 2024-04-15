package main

import (
	"fmt"

	"github.com/saifhatric/blockchain-v1/block"

	"github.com/saifhatric/blockchain-v1/wallet"
)

func main() {
	// bc := blockchain.NewBlockChain("my_Blockchain_Addr")

	// //Block 01
	// bc.AddTransactions("A", "B", 1.7)
	// bc.Mining()

	// //Block 02
	// bc.AddTransactions("C", "D", 3.2)
	// bc.AddTransactions("X", "Y", 3.14)
	// bc.Mining()

	// bc.Print()
	// fmt.Printf("my %.1f\n", bc.CalculateAmount("my_Blockchain_Addr"))
	// fmt.Printf("my %.1f\n", bc.CalculateAmount("A"))
	// fmt.Printf("my %.1f\n", bc.CalculateAmount("C"))

	// w := wallet.NewWallet()
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddr)
	// t := wallet.NewTransaction(w.PrivetKey, w.PublicKey, w.BlockchainAddr, "B", 2.4)
	// fmt.Printf("signature : %s\n", utils.GenerateSignature(t).SignatureStr())

	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// wallet side
	wt := wallet.NewTransaction(walletA.PrivetKey, walletA.PublicKey, walletA.BlockchainAddr, walletB.BlockchainAddr, 1.12)

	// service side (blockchain)
	bc := block.NewBlockChain(walletM.BlockchainAddr)

	// adding a transaction to the transaction pool
	isAdded := bc.AddTransactions(wt, walletA.BlockchainAddr, walletB.BlockchainAddr, wt.Value)

	fmt.Println("Added ?", isAdded)

	bc.Mining()
	bc.Print()

	fmt.Printf("A %.2f\n", bc.CalculateAmount(walletA.BlockchainAddr))
	fmt.Printf("B %.2f\n", bc.CalculateAmount(walletB.BlockchainAddr))
	fmt.Printf("M %.2f\n", bc.CalculateAmount(walletM.BlockchainAddr))

}
