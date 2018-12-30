package main

import (
	"github.com/boltdb/bolt"
)

const dbFile = "blockChainDb.db"
const blockBucket = "block"
const lastHashString = "lastHash"

/*
区块链
*/
type BlockChain struct {
	db       *bolt.DB
	lastHash []byte
}

//新建区块链
func NewBlockChain() *BlockChain {
	db, err := bolt.Open(dbFile, 0600, nil)
	var lastHash []byte
	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		if bucket != nil {
			//有数据
			lastHash = bucket.Get([]byte(lastHashString))
		} else {
			bucket, err := tx.CreateBucket([]byte(blockBucket))
			CheckErr("CreateBucket", err)
			//创建bucket,创建创始链
			lastHash = saveBlock(bucket, NewGenesisBlock())
		}
		return nil
	})
	CheckErr("NewBlockChain", err)
	return &BlockChain{
		db, lastHash,
	}
}
func saveBlock(bucket *bolt.Bucket, block *Block) []byte {
	err := bucket.Put(block.Hash, block.Serialize())
	CheckErr("saveBlock", err)
	_ = bucket.Put([]byte(lastHashString), block.Hash)
	CheckErr("saveBlock", err)
	return block.Hash
}

/**
添加区块
 */
func (bc *BlockChain) AddBlock(data string) {
	var prevBlockHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		prevBlockHash = bucket.Get([]byte(lastHashString))
		return nil
	})
	CheckErr("AddBlock", err)
	block := NewBlock(data, bc.lastHash)
	err = bc.db.Update(func(tx *bolt.Tx) error {
		bc.lastHash = saveBlock(tx.Bucket([]byte(blockBucket)), block)
		return nil
	})
	CheckErr("db.Update", err)
}

/**
区块链迭代器
 */
type BlockChainIterator struct {
	db          *bolt.DB
	currentHash []byte
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db, bc.lastHash}
}

func (it *BlockChainIterator) Next() *Block {
	var block *Block
	err := it.db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(blockBucket)).Get(it.currentHash)
		block = Deserialize(data)
		return nil
	})
	CheckErr("Next",err)
	it.currentHash = block.PrevBlockHash
	return block
}
