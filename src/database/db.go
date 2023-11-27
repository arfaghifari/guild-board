package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var (
		err error
	)

	cnxn := "postgres://naufal.ghifari:user@localhost:5432/guild?sslmode=disable"
	DB, err = sql.Open("postgres", cnxn)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func GetDB() *sql.DB {
	if DB == nil {
		InitDB()
	}
	return DB
}
