package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const dbhost = ""

var connect_db *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", DBUser+":"+DBPass+"@/"+DBName+"?parseTime=true")

	if err != nil {
		panic(err)
	} else {
		connect_db = db
		log.Println("Database berhasil terkoneksi!")
	}
}
