package utils

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// GetDynamicQuery is used to quickly test select statements by returning an array of structs
func GetDynamicQuery(s string) ([][]string, error) {
	var results [][]string


	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println("dbHost\ndbUser\ndbPassword\ndbDatabase\ndbPort", dbHost, dbUser)

	db, err := sql.Open("mysql", dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":"+ dbPort +")/" + dbDatabase + "?parseTime=true")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	rows, err := db.Query(s)

	if err != nil {
		return results, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return results, err
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice

	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return results, err
		}

		container := make([]string, len(cols))

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}

			container = append(container, result[i])
		}

		results = append(results, container)
	}

	return results, nil
}