package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/GeorgeGogos/Company-Management/controllers"
	"github.com/GeorgeGogos/Company-Management/middlewares"
)

func RegisterRoutes(r *mux.Router, companyController *controllers.CompanyController, authController *controllers.AuthController, userController *controllers.UserController) {
	// Company routes (require JWT)
	r.Handle("/companies", middlewares.JWTAuthMiddleware(http.HandlerFunc(companyController.CreateCompany))).Methods("POST")
	r.Handle("/companies/{id}", middlewares.JWTAuthMiddleware(http.HandlerFunc(companyController.UpdateCompany))).Methods("PATCH")
	r.Handle("/companies/{id}", middlewares.JWTAuthMiddleware(http.HandlerFunc(companyController.DeleteCompany))).Methods("DELETE")
	r.HandleFunc("/companies/{id}", companyController.GetCompany).Methods("GET")

	// Auth routes
	r.HandleFunc("/login", authController.Login).Methods("POST")
	r.HandleFunc("/register", userController.Register).Methods("POST")
}
