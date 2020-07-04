package services

import (
	"encoding/base64"
	"fmt"

	model "../../domain/entities"
	interfaces "../../domain/interfaces"
	repository "../../infraestructure/repositories"
	service "../interfaces"
)

type blockserv struct{}

var blockRepo interfaces.IBlockChainRepository = repository.BlockChainRepository()

// Datos struct
type Datos struct {
	Name     string
	Birthday string
	Dni      string
	Telefono string
}

// BlockChainService Implemantation
func BlockChainService() service.IBlockChainService {
	return &blockserv{}
}

func (*blockserv) BlockChainGenerated(block model.Block) (bool, error) {
	r, _ := blockRepo.GetAllBlocks()
	if len(r) == 0 {
		CreateBlockChain()
	}
	responseBlock, _ := blockRepo.GetLastBlock()
	block.Previous = responseBlock.Hash
	data := Datos{
		Name:     block.Name,
		Birthday: block.Birthday,
		Dni:      block.Dni,
		Telefono: block.Telefono,
	}
	block.Hash = GenerateHash(responseBlock.ID, data)
	_, err := blockRepo.AddNewBlock(block)
	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateBlockChain function
func CreateBlockChain() {
	block := model.Block{
		Previous: "0",
	}
	data := Datos{}
	block.Hash = GenerateHash(1, data)
	r := blockRepo.AddInitialBlock(block)
	if r != nil {
		fmt.Println("error: " + r.Error())
	}
}

// GenerateHash function
func GenerateHash(index int, data Datos) string {
	hash := fmt.Sprintf("%d-%s", index, data)
	return base64.StdEncoding.EncodeToString([]byte(hash))
}
