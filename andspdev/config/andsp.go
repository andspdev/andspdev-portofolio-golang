package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"andsp.id/andspdev/config/variables"
)

var connect_db *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", variables.DBUser+":"+variables.DBPass+"@/"+variables.DBName+"?parseTime=true")

	if err != nil {
		panic(err)
	} else {
		connect_db = db
		log.Println("Database berhasil terkoneksi!")
	}
}
