package main

import (
	"example.com/go-news-api/db"
	"example.com/go-news-api/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
