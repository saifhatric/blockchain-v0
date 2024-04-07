package main

import "github.com/saifhatric/blockchain-v1/blockchain"

func main() {
	bc := blockchain.NewBlockChain("my_Blockchain_Addr")

	//Block 01
	bc.AddTransactions("A", "B", 3.2)
	bc.Mining()

	//Block 02
	bc.AddTransactions("C", "D", 3.2)
	bc.AddTransactions("X", "Y", 3.14)
	bc.Mining()

	bc.Print()

}
