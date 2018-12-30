package main

import "fmt"

func main() {
	PrintBlock(Deserialize(NewGenesisBlock().Serialize()))
	//bc := NewBlockChain()
	//bc.AddBlock("第一次数据")
	//bc.AddBlock("第二次数据")
	//BlockChainIterator := bc.Iterator()
	//BlockChainIterator.PrintBlockChain()
}
func (b *BlockChainIterator) PrintBlockChain() {
	for b.currentHash != nil {
		PrintBlock(b.Next())
	}
}
func PrintBlock(block *Block) {
	fmt.Printf("block data :%s\n", block.d)
	fmt.Println("Version :", block.Version)
	fmt.Printf("PrevBlockHash :%x\n", block.PrevBlockHash)
	fmt.Printf("Hash :%x\n", block.Hash)
	fmt.Println("Timestamp :", block.TimeStamp)
	fmt.Println("MerKelRoot :", block.MerKelRoot)
	fmt.Println("Nonce :", block.Nonce)
	pow := NewProofOfWork(block)
	fmt.Println("Pow is valid :", pow.IsValid())
}
