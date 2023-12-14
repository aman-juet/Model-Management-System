package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection parameters
	dbUser := "root"
	dbPassword := "aman123@"
	dbName := "model"
	dbHost := "localhost"
	dbPort := "3306"

	// Create a connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Connected to the database successfully!")

	// Perform a simple query
	rows, err := db.Query("SELECT * FROM model.model")
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var column1 int
		err := rows.Scan(&column1)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Println("Column1:", column1)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}
}
