package persistance

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	con "../config"
	//connection
	_ "github.com/lib/pq"
)

var config con.Psql

// GetConnection obtiene una conexion a la BD
func GetConnection() *sql.DB {
	config.SetConfiguration()
	psql := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"))
	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
