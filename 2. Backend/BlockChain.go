package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// CovidRecord Struct
type CovidRecord struct {
	Name     string
	Birthday string
	Dni      string
	Telefono string
}

// Block Struct
type Block struct {
	Index        int
	Data         CovidRecord
	PreviousHash string
	Hash         string
}

// BlockChain Struct
type BlockChain struct {
	Chain []Block
}

// GenerateHash Function
func (block *Block) GenerateHash() string {
	hash := fmt.Sprintf("%d-%s", block.Index, block.Data)
	return base64.StdEncoding.EncodeToString([]byte(hash))
}

// GenerateInitialBlock Function
func (blockChain *BlockChain) GenerateInitialBlock() Block {
	block := Block{
		Index:        0,
		Data:         CovidRecord{},
		PreviousHash: "0",
	}
	block.Hash = block.GenerateHash()
	return block
}

// GetLastBlock Function
func (blockChain *BlockChain) GetLastBlock() Block {
	n := len(blockChain.Chain)
	return blockChain.Chain[n-1]
}

// AddBlock Function
func (blockChain *BlockChain) AddBlock(block Block) {
	block.Index = blockChain.GetLastBlock().Index + 1
	block.PreviousHash = blockChain.GetLastBlock().Hash
	block.Hash = block.GenerateHash()
	blockChain.Chain = append(blockChain.Chain, block)
}

// IsValid Function
func (blockChain *BlockChain) IsValid() bool {
	n := len(blockChain.Chain)
	for i := 1; i < n; i++ {
		current := blockChain.Chain[i]
		previous := blockChain.Chain[i-1]
		if current.Hash != current.GenerateHash() {
			fmt.Printf("current: %s y current actual: %s\n", current.Hash, current.GenerateHash())
			return false
		}
		if current.PreviousHash != previous.Hash {
			return false
		}
	}
	return true
}

// CreateBlockChain Function
func CreateBlockChain() BlockChain {
	blockChain := BlockChain{}
	initialBlock := blockChain.GenerateInitialBlock()
	blockChain.Chain = append(blockChain.Chain, initialBlock)
	return blockChain
}

func main() {
	blockChain := CreateBlockChain()
	covidRecord1 := CovidRecord{
		Name:     "Joaquin Gomez",
		Birthday: "21/07/1996",
		Dni:      "07566487",
		Telefono: "96437501",
	}
	covidRecord2 := CovidRecord{
		Name:     "Luis Mendoza",
		Birthday: "08/10/2001",
		Dni:      "48713246",
		Telefono: "95124789",
	}
	blockRecord1 := Block{
		Data: covidRecord1,
	}
	blockRecord2 := Block{
		Data: covidRecord2,
	}
	blockChain.AddBlock(blockRecord1)
	blockChain.AddBlock(blockRecord2)
	fmt.Println(blockChain.Chain[0])
	fmt.Println(blockChain.Chain[1])
	fmt.Println(blockChain.Chain[2])
	fmt.Printf("Is Valid?: %s\n", strconv.FormatBool(blockChain.IsValid()))

	//Se intenta alterar un Bloque
	blockChain.Chain[1].Data.Name = "Pedro Alvarez"
	fmt.Println(blockChain.Chain[0])
	fmt.Println(blockChain.Chain[1])
	fmt.Println(blockChain.Chain[2])
	fmt.Printf("Is Valid?: %s\n", strconv.FormatBool(blockChain.IsValid()))
}
