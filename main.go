package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	database "github.com/ordinary-fdev/go-learning/database"
	"github.com/ordinary-fdev/go-learning/router"
)

func main() {
	database.InitializeConnection()
	router.InitializeRoutes()
}
