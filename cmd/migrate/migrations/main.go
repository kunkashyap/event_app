package main

import (
	"fmt"
	"os"
	"github.com/golang-migrate/migrate/database/sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide migration directions : up or down")
		os.Exit(1)
	}

	direction := os.Args[1]

	db, err : sql.Open("sqlite3", './data.db')
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
}
