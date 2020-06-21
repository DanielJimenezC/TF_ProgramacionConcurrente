package persistance

import (
	"database/sql"
	"log"

	//connection
	_ "github.com/lib/pq"
)

// GetConnection obtiene una conexion a la BD
func GetConnection() *sql.DB {
	psql := "postgres://yvygupbmnbeljh:6fbd4ab6771e9fd2bb3bdbbda2a3e87b2c352f87ca3a8167d5b10c3c79af0cd1@ec2-52-72-221-20.compute-1.amazonaws.com:5432/dap059vc7a2ar2"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
