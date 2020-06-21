package entities

// User Model
type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

/*
// Create User Function
func Create(u User) error {
	q := `INSERT INTO public."user" (username, password) VALUES ($1, $2)`
	db := context.GetConnection()
	defer db.Close()

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(u.Username, u.Password)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()

	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}

	return nil
}*/
