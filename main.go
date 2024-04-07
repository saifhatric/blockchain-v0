package main

import (
	"fmt"

	"github.com/saifhatric/blockchain-v1/blockchain"
)

func main() {
	bc := blockchain.NewBlockChain("my_Blockchain_Addr")

	//Block 01
	bc.AddTransactions("A", "B", 1.7)
	bc.Mining()

	//Block 02
	bc.AddTransactions("C", "D", 3.2)
	bc.AddTransactions("X", "Y", 3.14)
	bc.Mining()

	bc.Print()
	fmt.Printf("my %.1f\n", bc.CalculateAmount("my_Blockchain_Addr"))
	fmt.Printf("my %.1f\n", bc.CalculateAmount("A"))
	fmt.Printf("my %.1f\n", bc.CalculateAmount("C"))

}
