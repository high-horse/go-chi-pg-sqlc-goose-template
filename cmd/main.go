package main

import (
	// "database/sql"
	"log"
	"net/http"
	"os"

	// "os"
	// "server/http/router"
	db "server/init"
	router"server/http/router"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func main() {
	err := godotenv.Load();
	if err != nil {
		log.Fatal("ERROR loading env :",err)
	}
	
	err = db.ConnectDB() 
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.DisconnectDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serve(port)

}


func serve(port string) {

	router := router.InitRouter()

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