// The utility helps to verify that you can
// connect to a PostgreSQL database and get
// a list of available databases and tables
// in that database and public schema.
package main

// Imports all of the needed packages.
import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Try to parse the arguments from the cmd .
	arguments := os.Args

	if len(arguments) != 6 {
		fmt.Println("дай - hostname post username password db")
		return
	}

	// Put the values to the vars.
	host := arguments[1]
	port := arguments[2]
	usrName := arguments[3]
	password := arguments[4]
	dbName := arguments[5]

	// Create string of connection data.
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, usrName, password, dbName)

	// Try to open connection to db.
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()

	// Try to read databases names.
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

	// Try to read tables names.
	rows, err = db.Query(`SELECT table_name FROM information_schema.tables
	WHERE table_schema='public'
	ORDER BY table_name`)

	if err != nil {
		fmt.Println("Query", err)
		return
	}

	// BUG(r): Just the bug example)).
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
