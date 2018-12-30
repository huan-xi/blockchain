package main

import (
	"time"
)

//区块
type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte //比特币中是不包含
	TimeStamp     int64
	TargetBits    int64
	data          []byte
	Nonce         int64
	MerKelRoot    []byte
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: preBlockHash,
		TimeStamp:     time.Now().Unix(),
		TargetBits:    targetBits,
		Nonce:         0,
		MerKelRoot:    []byte{},
		data:          []byte(data),
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

//创世块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
