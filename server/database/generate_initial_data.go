package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


func setupDB() *sql.DB {
    err := godotenv.Load("../../.env")
    var (
        DB_USER     = os.Getenv("DB_USER")
        DB_PASSWORD = os.Getenv("DB_PASSWORD")
        DB_NAME     = os.Getenv("DB_NAME")
    )
    checkErr(err)
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return db
}


func rowExists(query string, args ...interface{}) bool {
	db := setupDB()
    var exists bool
    query = fmt.Sprintf("SELECT exists (%s)", query)
    err := db.QueryRow(query, args...).Scan(&exists)
    if err != nil && err != sql.ErrNoRows {
        checkErr(err)
    }
    return exists
}


func main()  {
	db := setupDB()
	sqlStatement := `
		CREATE TABLE IF NOT EXISTS fuel_prices (
			id SERIAL PRIMARY KEY,
			qty INT NOT NULL,
			premium_price INT NOT NULL,
			pertalite_price INT NOT NULL
		);
	`
	_, err := db.Exec(sqlStatement)
	checkErr(err)
	for i := 1; i <= 20; i++ {
		qty := i
		premium_price := 6450 * i
		pertalite_price := 7650 * i

		if rowExists("SELECT id FROM fuel_prices WHERE qty=$1", qty) {
			continue;
		}

		var lastInsertID int
		db.QueryRow("INSERT INTO fuel_prices(qty, premium_price, pertalite_price) VALUES($1, $2, $3) returning id;", qty, premium_price, pertalite_price).Scan(&lastInsertID)
	}
}