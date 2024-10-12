package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/GeorgeGogos/Company-Management/config"
	"github.com/GeorgeGogos/Company-Management/controllers"
	"github.com/GeorgeGogos/Company-Management/repository"
	"github.com/GeorgeGogos/Company-Management/routes"
	"github.com/GeorgeGogos/Company-Management/services"
)

func main() {
	// Initialize DB
	db := config.ConnectDB()
	defer db.Close()

	// Initialize repositories, services, and controllers
	companyRepo := &repository.CompanyRepository{DB: db}
	userRepo := &repository.UserRepository{DB: db}

	companyService := &services.CompanyService{Repo: companyRepo}
	companyController := &controllers.CompanyController{Service: companyService}
	authController := &controllers.AuthController{UserRepo: userRepo}
	userController := &controllers.UserController{UserRepo: userRepo}

	// Initialize router and register routes
	r := mux.NewRouter()
	routes.RegisterRoutes(r, companyController, authController, userController)

	// Start the server
	log.Println("Starting server on port 30000")
	log.Fatal(http.ListenAndServe(":30000", r))
}