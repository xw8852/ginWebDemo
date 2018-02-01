package database

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Database struct {
	Db *sql.DB
}

func Default() *Database {
	db, err := sql.Open("mysql", "miniapp:mini8875901@tcp(47.94.110.244:3306)/miwei?charset=utf8")
	if (err != nil) {
		log.Fatal(err)
	}
	var database = &Database{}
	database.Db = db
	return database
}
