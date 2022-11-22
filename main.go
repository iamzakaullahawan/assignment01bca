package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	transaction  string
	nonce        int
	previousHash string
	blockHash    string
}

func (blk *Block) CalculateHash(stringToHash string) {
	sum := sha256.Sum256([]byte(stringToHash))
	Hash := hex.EncodeToString(sum[:])
	blk.blockHash = Hash
}

type Chain struct {
	blockchain []Block
	chain_hash string
}

func (bc1 *Chain) GetChainHash() string {
	return bc1.chain_hash
}

func (bc1 *Chain) NewBlock(transaction string, nonce int, previousHash string) *Block {
	var new_block = Block{transaction: transaction, nonce: nonce, previousHash: previousHash}
	new_block.CalculateHash(transaction + strconv.Itoa(nonce) + previousHash)
	bc1.blockchain = append(bc1.blockchain, new_block)
	bc1.chain_hash = bc1.blockchain[len(bc1.blockchain)-1].blockHash
	return &new_block
}

func (bc1 *Chain) ListBlocks() {
	for i := range bc1.blockchain {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
		fmt.Printf("transaction: %s\n", bc1.blockchain[i].transaction)
		fmt.Printf("nonce: %d\n", bc1.blockchain[i].nonce)
		fmt.Printf("Previous Hash: %s\n", bc1.blockchain[i].previousHash)
		fmt.Printf("Block Hash: %s\n\n", bc1.blockchain[i].blockHash)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 58))
	fmt.Printf("Chain Hash: %s\n", bc1.chain_hash)
	fmt.Printf("%s\n", strings.Repeat("=", 58))
}

func (bc1 *Chain) ChangeBlock() {
	bc1.blockchain[1].transaction = "This transcation has been changed."
	bc1.blockchain[1].CalculateHash(bc1.blockchain[1].transaction + strconv.Itoa(bc1.blockchain[1].nonce) + bc1.blockchain[1].previousHash)
}

func (bc1 *Chain) VerifyChain() {
	bc1.blockchain[0].CalculateHash(bc1.blockchain[0].transaction + strconv.Itoa(bc1.blockchain[0].nonce) + "0")

	for i := 1; i < len(bc1.blockchain); i++ {
		bc1.blockchain[i].CalculateHash(bc1.blockchain[i].transaction + strconv.Itoa(bc1.blockchain[i].nonce) + bc1.blockchain[i-1].blockHash)
	}

	fmt.Printf("Previous Chain Hash: %s \n", bc1.chain_hash)
	fmt.Printf("New Chain Hash: %s \n\n", bc1.blockchain[len(bc1.blockchain)-1].blockHash)

	if bc1.chain_hash == bc1.blockchain[len(bc1.blockchain)-1].blockHash {
		fmt.Printf("Block Chain is not modified. \n")

	} else {
		fmt.Printf("Block Chain is modified.\n")

	}
	fmt.Printf("%s\n\n\n", strings.Repeat("=", 58))

}
