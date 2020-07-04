package repositories

import (
	"errors"
	"fmt"

	model "../../domain/entities"
	interfaces "../../domain/interfaces"
	context "../persistance"
)

type repoBlock struct{}

// BlockChainRepository Implement
func BlockChainRepository() interfaces.IBlockChainRepository {
	return &repoBlock{}
}

func (*repoBlock) AddInitialBlock(block model.Block) error {
	q := `INSERT INTO public."block" (previous, hash) VALUES ($1, $2)`
	db := context.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		fmt.Println("aca: " + err.Error())
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(block.Previous, block.Hash)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}

func (*repoBlock) AddNewBlock(block model.Block) (bool, error) {
	q := `INSERT INTO public."block" (name, birthday, dni, telefono, previous, hash) VALUES ($1, $2, $3, $4, $5, $6)`
	db := context.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	r, err := stmt.Exec(block.Name, block.Birthday, block.Dni, block.Telefono, block.Previous, block.Hash)
	if err != nil {
		return false, err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return false, errors.New("Error: se esperaba 1 fila afectada")
	}

	return true, nil
}

func (*repoBlock) GetLastBlock() (model.Block, error) {
	q := `SELECT id, hash FROM public."block" ORDER BY id DESC LIMIT 1`
	db := context.GetConnection()
	defer db.Close()

	var block model.Block

	r, err := db.Query(q)
	if err != nil {
		return block, err
	}
	defer r.Close()

	for r.Next() {
		r.Scan(&block.ID, &block.Hash)
		if err != nil {
			return block, err
		}
	}
	return block, nil
}

func (*repoBlock) GetAllBlocks() ([]model.Block, error) {
	q := `SELECT * FROM public."block"`
	db := context.GetConnection()
	defer db.Close()

	r, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var blocks []model.Block

	for r.Next() {
		e := model.Block{}
		r.Scan(&e.ID, &e.Name, &e.Birthday, &e.Dni, &e.Telefono, &e.Previous, &e.Hash)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, e)
	}
	return blocks, nil
}
