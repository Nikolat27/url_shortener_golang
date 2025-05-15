package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"url_shortener/internal/routes"
)

func main() {
	server := &http.Server{
		Addr:    ":8000",
		Handler: routes.Routes(),
	}

	dsn := os.Getenv("DB_DSN")
	db, err := initDB(dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return
	}
	
	
	err = server.ListenAndServe()
	if err != nil {
		log.Println("server listening error: ", err)
	}
}

func initDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return db, err
}
