package routes

import (
	"net/http"

	"github.com/GeorgeGogos/Company-Management/controllers"
	"github.com/GeorgeGogos/Company-Management/middlewares"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, controller *controllers.CompanyController) {
	r.Handle("/companies", middlewares.JWTAuthMiddleware(http.HandlerFunc(controller.CreateCompany))).Methods("POST")
	r.Handle("/companies/{id}", middlewares.JWTAuthMiddleware(http.HandlerFunc(controller.UpdateCompany))).Methods("PATCH")
	r.Handle("/companies/{id}", middlewares.JWTAuthMiddleware(http.HandlerFunc(controller.DeleteCompany))).Methods("DELETE")
	r.HandleFunc("/companies/{id}", controller.GetCompany).Methods("GET")
}
