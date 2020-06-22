package repositories

import (
	"errors"

	model "../../domain/entities"
	interfaces "../../domain/interfaces"
	context "../persistance"
)

type repo struct{}

// UserRepository Implement
func UserRepository() interfaces.IUserRepository {
	return &repo{}
}

func (*repo) Create(user model.User) error {
	q := `INSERT INTO public."user" (username, password) VALUES ($1, $2)`
	db := context.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}

func (*repo) GetAll() ([]model.User, error) {
	q := `SELECT * FROM public."user" ORDER BY id ASC`
	db := context.GetConnection()
	defer db.Close()

	r, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var users []model.User

	for r.Next() {
		e := model.User{}
		r.Scan(&e.ID, &e.Username, &e.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, e)
	}
	return users, nil
}

func (*repo) GetUser(id int) (model.User, error) {
	q := `SELECT * FROM public."user" WHERE id = $1`
	db := context.GetConnection()

	var user model.User
	row := db.QueryRow(q, id)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, nil
	}
	return user, nil
}

func (*repo) Delete(id int) error {
	q := `DELETE FROM public."user" WHERE id = $1`
	db := context.GetConnection()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}

func (*repo) Update(id int, user model.User) error {
	q := `UPDATE public."user" SET username=$1, password=$2 WHERE id = $3`
	db := context.GetConnection()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(user.Username, user.Password, id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}
