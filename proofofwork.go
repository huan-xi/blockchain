package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 10

/**
工作量证明：
根据nonce不同对不同的数据取hash值直到符合设定的难度值
*/
type ProofOfWork struct {
	block     *Block
	targetBit *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {
	var IntTarget = big.NewInt(1)
	//左移
	IntTarget.Lsh(IntTarget, uint(256-targetBits))
	return &ProofOfWork{
		block,
		IntTarget,
	}
}

func (pow *ProofOfWork) PrepareRawData(nonce int64) []byte {
	block := pow.block
	tmp := [][]byte{
		IntToByte(block.Version),
		block.PrevBlockHash,
		IntToByte(block.TimeStamp),
		IntToByte(nonce),
		IntToByte(targetBits),
		block.data,
	}
	data := bytes.Join(tmp, []byte{})
	return data
}

//碰撞计算hash
func (pow *ProofOfWork) Run() (int64, []byte) {
	var nonce int64
	var hash [32]byte
	var hashInt big.Int
	fmt.Println("开始挖矿........")
	for nonce < math.MaxInt64 {
		data := pow.PrepareRawData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.targetBit) == -1 {
			fmt.Printf("找到Hash:%x\n", hash)
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) IsValid() bool {
	data := pow.PrepareRawData(pow.block.Nonce)
	var HashInt big.Int
	hash := sha256.Sum256(data)
	HashInt.SetBytes(hash[:])
	return HashInt.Cmp(pow.targetBit) == -1
}
