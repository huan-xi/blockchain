package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("第一次数据")
	bc.AddBlock("第二次数据")
	for i, block := range bc.blocks {
		fmt.Println("============block num :",i)
		fmt.Printf("block data :%s\n",block.data)
		fmt.Println("Version :",block.Version)
		fmt.Printf("PrevBlockHash :%x\n",block.PrevBlockHash)
		fmt.Printf("Hash :%x\n",block.Hash)
		fmt.Println("Timestamp :",block.TimeStamp)
		fmt.Println("MerKelRoot :",block.MerKelRoot)
		fmt.Println("Nonce :",block.Nonce)
		pow:=NewProofOfWork(block)
		fmt.Println("Pow is valid :",pow.IsValid())
	}
}
