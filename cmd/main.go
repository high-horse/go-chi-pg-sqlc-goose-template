package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"server/sql/database"
	"server/router"
	"server/models"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


type DBConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("ERROR loading env :",err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("ERROR connectiong to database :",err)
	}

	db := database.New(conn)
	dbConfig  := models.DBConfig{
		DB: db,
	}

	serve(dbConfig, port)

}


func serve(dbConfig models.DBConfig, port string) {

	router := router.InitRouter(dbConfig)

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}
	println("Server running on port", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ERROR starting server :",err)
	}
}