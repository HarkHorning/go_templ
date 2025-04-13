package database

import (
	"database/sql"
	"fmt"
	"log"
	// "github.com/go-sql-driver/mysql"
)

func main() {

	// Connection string
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname"
	// Open a database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")

}
