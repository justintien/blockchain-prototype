package main

type BlockChain struct {
	Chain []*Block
}

func (bc *BlockChain) CreateGenesisBlock() *Block {
	b := &Block{
		Timestamp:     "2018-01-01",
		Data:          "Genesis Block",
		PrevBlockHash: "0",
	}
	b.SetHash()
	return b
}

func (bc *BlockChain) GetLatestBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *BlockChain) AddBlock(newBlock *Block) {
	newBlock.PrevBlockHash = bc.GetLatestBlock().Hash
	newBlock.SetHash()
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *BlockChain) IsChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]
		currentBlock.SetHash()
		if currentBlock.Hash != currentBlock.Hash {
			return false
		}

		if currentBlock.PrevBlockHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
