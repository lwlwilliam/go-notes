// 区块链操作
package core

import (
	"log"
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

// 创建区块链
func NewBlockchain() *Blockchain {
	// 先生成创世区块
	genesisBlock := GenerateGenesisBlock()

	// 创建区块链并把创世区块添加到其中
	blockchain := Blockchain{}
	blockchain.AppendBlock(&genesisBlock)
	return &blockchain
}

// 生成新区块并往区块链中添加该区块
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

// 添加区块到区块链中
func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

// 打印所有区块的信息
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev.Hash: %s\n", block.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Println()
	}
}

// 判断新增的区块是否合法
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}

	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}
