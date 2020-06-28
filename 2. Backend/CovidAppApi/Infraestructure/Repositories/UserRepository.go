package repositories

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

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

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	r, err := stmt.Exec(user.Username, string(hash))
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
	defer db.Close()
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
	defer db.Close()

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
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	if user.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		r, err := stmt.Exec(user.Username, string(hash), id)
		if err != nil {
			return err
		}
		i, _ := r.RowsAffected()

		if i != 1 {
			return errors.New("Error: Se esperaba 1 fila afectada")
		}
	} else {
		r, err := stmt.Exec(user.Username, user.Password, id)
		if err != nil {
			return err
		}
		i, _ := r.RowsAffected()

		if i != 1 {
			return errors.New("Error: Se esperaba 1 fila afectada")
		}
	}

	return nil
}

func (*repo) Login(user model.User) (bool, error) {
	q := `SELECT * FROM public."user" WHERE username=$1`
	db := context.GetConnection()
	defer db.Close()

	var loginUser model.User
	row := db.QueryRow(q, user.Username)
	var hash []byte
	err := row.Scan(&loginUser.ID, &loginUser.Username, &loginUser.Password)
	if err != nil {
		return false, nil
	}
	hash = []byte(loginUser.Password)
	password := user.Password
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return false, nil
	}

	os.Setenv("USER_LOGIN", strconv.Itoa(loginUser.ID))
	return true, nil
}
