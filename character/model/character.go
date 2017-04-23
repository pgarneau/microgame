package model

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Character struct {
	Name  string `json:"name"`
	Class string `json:"class"`
	Level int    `json:"level"`
	ID    int    `json:"id"`
}

type Characters []Character

var db *sql.DB

func ListDB() *sql.DB {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=mypass host=db port=5432 sslmode=disable")
	if err != nil {
		fmt.Printf("waddupss")
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("O SHIIIT")
	}

	_, err = db.Exec(
		`INSERT INTO characters(name, class) VALUES ('Phil', 'Paladin')`)

	if err != nil {
		fmt.Printf("FAILED TABLE")
	}

	return db
}
