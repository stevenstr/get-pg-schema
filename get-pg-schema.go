package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	arguments := os.Args

	if len(arguments) != 6 {
		fmt.Println("дай - hostname post username password db")
		return
	}

	host := arguments[1]
	port := arguments[2]
	usrName := arguments[3]
	password := arguments[4]
	dbName := arguments[5]

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, usrName, password, dbName)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "datname" FROM "pg_database"
	WHERE datistemplate = false`)

	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan:", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()

	rows, err = db.Query(`SELECT table_name FROM information_schema.tables
	WHERE table_schema='public'
	ORDER BY table_name`)

	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan:", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()

}
