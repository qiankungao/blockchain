package main

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevblock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevblock.PrevBlockHash)

	bc.blocks = append(bc.blocks, newBlock)

}
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
