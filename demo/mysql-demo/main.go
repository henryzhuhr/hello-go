package main

import "database/sql"

func main() {

	database, err := sql.Open("mysql", "root:password@tcp()")
	if err != nil {
		panic(err)
	}
	defer database.Close()
}
