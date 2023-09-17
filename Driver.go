package main

import (
	"fmt"
	"os"

	"github.com/mhuzaifi0604/assignment01bca"
)

// Driver Function for the block Chain
func main() {
	// Creating object of BlockChain list struct
	blockchain := &assignment01bca.Blockchain{}
	var input string
	var choice int
	// Creating a Genesis block in the block Chain
	assignment01bca.CreateAndInsert(0, "Genesis Block", blockchain)
	fmt.Println("Genesis Block Created on its own!")
	fmt.Println("")
	// Asking user a choice for addind or viewing the blockchain
	for {
		fmt.Printf("Choose an option:\n1- Add a new block\n2- View the BlockChain\n3- Change a Block\n4- Verify BlockChain\nChoose: ")
		fmt.Scanln(&choice)
		// Adding a new block to the chain
		if choice == 1 {
			// getting data from user for the Chain
			fmt.Print("Enter Data for new Block : ")
			fmt.Scanln(&input)
			// Creating a new block and inserting it in the Chain
			assignment01bca.CreateAndInsert(1, input, blockchain)
		} else if choice == 2 { // Printing the Whole BlockChain
			fmt.Printf("\t\t\t\t ============== Printing the Block Chain ============== \n\n")
			assignment01bca.PrintChain(blockchain)
		} else if choice == 3 {
			var new string
			var old string
			fmt.Print("Enter Data of the Block to Change: ")
			fmt.Scanln(&old)
			fmt.Print("Enter New Data to replace with: ")
			fmt.Scanln(&new)
			if assignment01bca.Change_block(old, new, blockchain) {
				fmt.Println("Changes made to the respective block in the Chain!")
			} else {
				fmt.Println("No Such Block found in the Chain!")
			}
		} else if choice == 4 { // Verifying the integrity of Chain
			if assignment01bca.Verify_BlockChain(blockchain) { // if All prev hashes match the hashes of prev block's current hashes
				fmt.Println("\nIntegrity of the Blockchain is Verified\n")
			} else { // If there's been a change in any of the Blocks
				fmt.Println("\nTheres been a change in one of the blocks. blockchain's Integrity Has been compromised!\n")
			}
		} else if choice == 0 { // Default invalid input
			fmt.Println("\nGood Bye!.")
			os.Exit(0)
		}
	}
}
