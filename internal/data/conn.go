package data

import (
	"database/sql"
	"easypwn/assets/mock"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func init() {
	dbUrl, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Println("DATABASE_URL is not set, running in test mode")

		var err error
		db, err = sql.Open("sqlite3", ":memory:")
		if err != nil {
			log.Fatal(err)
		}

		mockSql := mock.InitdbMockSql

		_, err = db.Exec(string(mockSql))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	var err error
	db, err = sql.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
