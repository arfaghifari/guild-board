package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "127.0.0.1"
// 	port     = 5432
// 	user     = "naufal.ghifari"
// 	password = ""
// 	dbname   = "guild"
// )

var DB *sql.DB

func InitDB() {
	var (
		err error
	)

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
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
