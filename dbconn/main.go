package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT host, user FROM mysql.user")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()

	var host, uname string

	for rows.Next() {
		err = rows.Scan(&host, &uname)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(host, uname)
	}

}
