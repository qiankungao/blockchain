package main

import (
	"github.com/boltdb/bolt"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	var tip []byte

	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *blot.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			gensis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))

			err = b.put(gensis.Hash, gensis.Serialize())
			err = b.put([]byte("1"), gensis.Hash)
			tip = gensis.Hash
		} else {
			tip = b.Get([]byte("1"))
		}
		return nil
	})
	bc := Blockchain{tip, db}
	return &bc
}
