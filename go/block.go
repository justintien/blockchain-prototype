package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Timestamp     string
	Data          string
	PrevBlockHash string
	Hash          string
}

func (b *Block) SetHash() {
	s := fmt.Sprintf("%s%s%s", b.PrevBlockHash, b.Timestamp, b.Data)
	sum := sha256.Sum256([]byte(s))

	b.Hash = fmt.Sprintf("%x", sum)
}
