package assignment01bca

// Importing Necessory Libraries
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

// Making a struct to contain attributes of a Block in Blockchain
type Block struct {
	Nonce        int
	Data         string
	CurrentHash  string
	PreviousHash string
}

// Creating a struct for list of above Blocks to create a Block Chain
type Blockchain struct {
	Chain []Block
}

// Funciton to insert New Block in the BlockChain
// Function takes all attributes of Struct Block to add new block into the Chain
func (chain *Blockchain) InsertBlock(nonce int, data string, previousHash string) {
	// Creating a new block using Block info recieved as parameters
	newBlock := Block{
		Nonce:        nonce,
		Data:         data,
		PreviousHash: previousHash,
		CurrentHash:  CalculateBlockHash(nonce, data, previousHash), // Calculating Block Hah
	}
	// Appending the new block in the Chain
	chain.Chain = append(chain.Chain, newBlock)
	fmt.Println("\nBlock Added To the Blockchain\n")
	// Printing Block Details After its been added to the Chain
	//fmt.Printf("Block Details: %+v\n", newBlock)
}

// Calculating a Nonce for the block to be added
func CalculateNonce(data string, previousHash string) int {
	var nonce int = 0
	for {
		// Calculating hash using nonce until desired hash is recieved
		hash := CalculateBlockHash(nonce, data, previousHash)
		// checking if hash calculated with current nide starts with 4 0s
		if hash[:4] == "0000" {
			fmt.Printf("\nNonce Calculated: %d\n", nonce) // Printing selected nonce and returning
			return nonce
		}
		nonce++
	}
}

// Function for calculating the hash of the block itself
// Taking all attributes except current hash as parameters to calculate current hashh
func CalculateBlockHash(nonce int, data string, previousHash string) string {
	// Accumulating all attributes into a single string
	BlockData := fmt.Sprintf("%d%s%s", nonce, data, previousHash)
	BlockHash := sha256.Sum256([]byte(BlockData)) // Calculating hash of block attributes
	return hex.EncodeToString(BlockHash[:])       // Returning the hash in string format
}

// Function for creating and inserting Blocks in to blockchain
// Recieving block type to add, data and Blockchain's object as parameters
func CreateAndInsert(blockType int, data string, object *Blockchain) {
	// Creating Genesis Block
	if blockType == 0 {
		gen_hash := strings.Repeat("0", 64)                                // Hash of Genesis Block is all Zeros
		object.InsertBlock(CalculateNonce(data, gen_hash), data, gen_hash) // Inserting Genesis Block into the blockchain
	} else if blockType == 1 { // Inserting Normal Block into the blockchain
		if len(object.Chain) > 0 { // Checking if Chain has a Genesis Block or not
			// Inserting new block in BlockChain after Genesis Block
			object.InsertBlock(CalculateNonce(data, object.Chain[len(object.Chain)-1].CurrentHash), data, object.Chain[len(object.Chain)-1].CurrentHash)
		} else { // Printing that since no Genesis block exists normal block can't be added
			fmt.Println("No previous block exists. Please create a genesis block first.")
		}
	}
}

// function to verify the integrity of the chain
func Verify_BlockChain(chain *Blockchain) bool {
	fmt.Println()
	// running loop from 1 since prev hash of block 0 (Genisis block) is 0
	for i := 1; i < len(chain.Chain); i++ {
		// Checking if prev hash of block matches the current hash if its prev block
		if chain.Chain[i].PreviousHash == chain.Chain[i-1].CurrentHash {
			fmt.Printf("Block [%d]'s Previous Hash Matches Block [%d]'s Current Hash\n", i, i-1)
		} else { // If it does not matches means that Block has been altered
			fmt.Printf("Block [%d]'s Previous Hash does not Match Block [%d]'s Current Hash\n", i, i-1)
			return false
		}
	}
	return true
}

// function to make a change in the blockchain
// Taking old data , new data and blockchain object as parameters
func Change_block(old_data string, new_data string, chain *Blockchain) bool {
	// traversing the chain to find block whose data == old_data
	for i := 0; i < len(chain.Chain); i++ {
		if chain.Chain[i].Data == old_data {
			// Updating the Block
			chain.Chain[i].Data = new_data
			chain.Chain[i].Nonce = CalculateNonce(chain.Chain[i].Data, chain.Chain[i].PreviousHash) // Calculating new nonce for the block
			// Calculating new current hash for the block
			chain.Chain[i].CurrentHash = CalculateBlockHash(chain.Chain[i].Nonce, chain.Chain[i].Data, chain.Chain[i].PreviousHash)
			return true
		}
	}
	return false
}

// Function  to print BlockChain
func PrintChain(chain *Blockchain) {
	for i := 0; i < len(chain.Chain); i++ {
		fmt.Printf("Block # %d: %v\n", i, chain.Chain[i]) // Priniting each block in the chain
	}
}
