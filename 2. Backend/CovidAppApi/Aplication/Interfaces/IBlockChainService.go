package services

import (
	entity "../../domain/entities"
)

// IBlockChainService Interface
type IBlockChainService interface {
	BlockChainGenerated(block entity.Block) (bool, error)
}
