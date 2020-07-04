package repositories

import entity "../entities"

// IBlockChainRepository Interface
type IBlockChainRepository interface {
	AddInitialBlock(block entity.Block) error
	AddNewBlock(block entity.Block) (bool, error)
	GetLastBlock() (entity.Block, error)
	GetAllBlocks() ([]entity.Block, error)
}
