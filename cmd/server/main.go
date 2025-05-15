package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"url_shortener/internal/api"
	"url_shortener/internal/data"
	"url_shortener/internal/routes"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	db, err := initDB(dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	models := data.NewModels(db)
	app := api.Application{
		Models: models,
	}

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes.Routes(&app),
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
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
