package main

import (
	"TarefasCalendario/models"
	"TarefasCalendario/routes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.CreateTasksTable()

	r := chi.NewRouter()
	routes.LoadRoutes(r)

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	address := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Handler: r,
		Addr:    address,
	}

	fmt.Printf("Servidor rodando em %v", address)
	log.Fatal(srv.ListenAndServe())

}
