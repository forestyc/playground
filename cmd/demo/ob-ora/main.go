package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func main() {
	var user, password, host, sqlCmd string
	flag.StringVar(&user, "u", "root", "user")
	flag.StringVar(&password, "p", "root", "password")
	flag.StringVar(&host, "h", "127.0.0.1:3306", "host")
	flag.StringVar(&sqlCmd, "s", "", "sql")
	flag.Parse()
	// Set up connection string
	dsn := user + "/" + password + "@" + host // Replace with your actual values

	// Open connection
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal("Failed to open connection:", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	fmt.Println("Successfully connected to OceanBase Oracle mode!")

	// Query example
	rows, err := db.Query(sqlCmd) // Replace with your table
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var column1 string
		if err := rows.Scan(&column1); err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		fmt.Println(column1)
	}
}
