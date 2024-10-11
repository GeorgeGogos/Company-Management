package main

import (
	"log"
	"net/http"

	"github.com/GeorgeGogos/Company-Management/config"
	"github.com/GeorgeGogos/Company-Management/controllers"
	"github.com/GeorgeGogos/Company-Management/repository"
	"github.com/GeorgeGogos/Company-Management/routes"
	"github.com/gorilla/mux"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	repo := &repository.CompanyRepository{DB: db}
	controller := &controllers.CompanyController{Repo: repo}

	r := mux.NewRouter()

	routes.RegisterRoutes(r, controller)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
