package main

import (
	"encoding/json"
	"fmt"
)

func NewBlockChain() *BlockChain {
	bc := &BlockChain{}
	bc.Chain = append(bc.Chain, bc.CreateGenesisBlock())
	return bc
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock(&Block{Timestamp: "2018-02-01", Data: `{"amout":10}`})
	bc.AddBlock(&Block{Timestamp: "2018-03-01", Data: `{"amout":20}`})
	bc.AddBlock(&Block{Timestamp: "2018-04-01", Data: `{"amout":30}`})

	blockchain, _ := json.MarshalIndent(bc, "", "\t")
	fmt.Println(string(blockchain))

	fmt.Println("is BlockChain valid ", bc.IsChainValid())
}
