package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeGogos/Company-Management/models"
	"github.com/GeorgeGogos/Company-Management/services"
	"github.com/gorilla/mux"
)

type CompanyController struct {
	Service *services.CompanyService
}

// CreateCompany handles the creation of a new company
func (c *CompanyController) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = c.Service.Create(&company)
	if err != nil {
		http.Error(w, "Unable to create company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

// GetCompany retrieves a company by its ID
func (c *CompanyController) GetCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	company, err := c.Service.Get(params["id"])
	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(company)
}

// UpdateCompany handles updating an existing company
func (c *CompanyController) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var company models.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	company.ID = params["id"]

	err = c.Service.Update(&company)
	if err != nil {
		http.Error(w, "Unable to update company", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(company)
}

// DeleteCompany handles deleting an existing company
func (c *CompanyController) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := c.Service.Delete(params["id"])
	if err != nil {
		http.Error(w, "Unable to delete company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}