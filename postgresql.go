// connect postgresql
package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "internDemo"
)

var psqlConnectString = fmt.Sprintf("host=%s port=%d user=%s "+
"password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)

func InitDB() (*sql.DB, error) {
    db, err := sql.Open("postgres", psqlConnectString)
    if err != nil {
        panic(err)
    }

	Ping(db)

	fmt.Println("Connected to PostgreSQL")
	return db, nil
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func Ping(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}
