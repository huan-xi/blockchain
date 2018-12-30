package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlock("第一次数据")
	bc.AddBlock("第二次数据")
	for i, block := range bc.blocks {
		fmt.Println("============block num :",i)
		fmt.Println("============block data :",block.data)
		fmt.Println("Version :",block.Version)
		fmt.Println("prevBlockHash :",block.PrevBlockHash)
	}
}
