package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const dbhost = ""
const dbuser string = "root"
const dbpass string = ""
const dbname string = "andsp_go"

const PortServer string = ":30902"
const SubRoutes string = "/andspdev"

var connect_db *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@/"+dbname+"?parseTime=true")

	if err != nil {
		panic(err)
	} else {
		connect_db = db
		log.Println("Database berhasil terkoneksi!")
	}
}
