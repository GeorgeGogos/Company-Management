package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeGogos/Company-Management/models"
	"github.com/GeorgeGogos/Company-Management/repository"
	"github.com/gorilla/mux"
)

type CompanyController struct {
	Repo *repository.CompanyRepository
}

func (c *CompanyController) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	_ = json.NewDecoder(r.Body).Decode(&company)

	err := c.Repo.Create(&company)
	if err != nil {
		http.Error(w, "Unable to create company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(company)
}

func (c *CompanyController) GetCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	company, err := c.Repo.Get(params["id"])

	if err != nil {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(company)
}

func (c *CompanyController) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var company models.Company
	_ = json.NewDecoder(r.Body).Decode(&company)
	company.ID = params["id"]

	err := c.Repo.Update(&company)
	if err != nil {
		http.Error(w, "Unable to update company", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(company)
}

func (c *CompanyController) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := c.Repo.Delete(params["id"])

	if err != nil {
		http.Error(w, "Unable to delete company", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
