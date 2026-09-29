package main

import (
	"database/sql"
	"fmt"
	"log"
	"os" //used for interacting with the operating system, such as reading command-line arguments and exiting the program

	"github.com/golang-migrate/migrate/database/sqlite3" //it allow to use SQLite3 as the database for migrations
)

func main() {
	if len(os.Args) < 2 { //checks whether user provided migration direction (up or down) as a command-line argument. If not, it prints an error message and exits the program.
		fmt.Println("Please provide migration directions : up or down")
		os.Exit(1)
	}

	direction := os.Args[1]

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", "fSrc", "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}
}
