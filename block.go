package main

import (
	"bytes"
	"crypto/sha256"
	//	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //当前时间戳
	Data          []byte //区块存储的实际有效信息
	PrevBlockHash []byte //前一个块的hash
	Hash          []byte //当前的hash
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})

	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
