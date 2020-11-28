package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://postgres:password@localhost/ecomgo?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("You are connected to database")
}
