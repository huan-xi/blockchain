package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//区块
type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte //正常是不包含
	TimeStamp     int64
	TargetBits    int64
	data          []byte
	Nonce         int64
	MerKelRoot    []byte
}

func (block *Block) SetHash() {
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		block.MerKelRoot,
		IntToByte(block.Nonce),
	}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: preBlockHash,
		TimeStamp:     time.Now().Unix(),
		TargetBits:    10,
		Nonce:         5,
		MerKelRoot:    []byte{},
		data:          []byte(data),
	}
	block.SetHash()
	return block
}
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}
