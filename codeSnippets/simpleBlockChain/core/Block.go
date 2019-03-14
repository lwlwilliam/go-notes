// 定义及生成区块
package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index int64				// 区块编号
	Timestamp int64			// 区块时间戳
	PrevBlockHash string	// 上一个区块哈希值
	Hash string				// 当前区块哈希值

	Data string				// 区块数据
}

// 计算区块的哈希值
func calculateHash(b Block) string {
	// 区块的哈希值与 index、timestamp、上一区块的 hash 和区块数据有关
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// 生成新的区块，需要传入上一区块以及新区块的值。其实也就使用上一区块的编号及 hash
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	// 区块编号是递增的
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

// 生成创世区块，也就是初始区块
func GenerateGenesisBlock() Block {
	// 创世区块的上一区块编号为 -1，那么创建区块自身的编号即为 0。创建区块的上一区块 hash 为空字符串
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}

