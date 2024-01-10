package database

import (
	"database/sql"
	"fmt"

	cfg "github.com/ordinary-fdev/go-learning/config"
)

var (
	db *sql.DB
)

func InitializeConnection() {
	config := cfg.ReadConfig()
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", config.Server, config.User, config.Pwd, config.Port, config.Db)
	conn, err := sql.Open("mssql", connString)

	if err != nil {
		panic(err)
	}
	db = conn
	fmt.Println("Connected with DB")
}

func GetDb() *sql.DB {
	return db
}
