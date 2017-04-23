package main

import (
	"database/sql"
	"fmt"
	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func main() {
	log.Printf("Connecting to DB")
	db, err := sql.Open("postgres", "user=postgres password=mypass host=db port=5432 sslmode=disable")
	if err != nil {
		fmt.Printf("PATATE")
	}

	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(i) * time.Second)
		log.Printf("WAITING...")

		if err = db.Ping(); err == nil {
			break
		}
		log.Println(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	migrator, _ := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "./go/postgresql")

	fmt.Printf("TRYING TO MIGRATE")
	err = migrator.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DOIIING STUFF")

}
