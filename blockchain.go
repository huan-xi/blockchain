package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		[]*Block{NewGenesisBlock()},
	}
}
func (bc *BlockChain) AddBlock(data string) {
	lastBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(data, lastBlock.Hash)
	bc.blocks = append(bc.blocks, block)
}
