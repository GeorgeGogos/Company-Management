package repository

import (
	"database/sql"
	"errors"

	"github.com/GeorgeGogos/Company-Management/models"
	"github.com/google/uuid"
)

type CompanyRepository struct {
	DB *sql.DB
}

func (r *CompanyRepository) Create(company *models.Company) error {
	company.ID = uuid.New().String()
	query := `INSERT INTO companies (id, name, description, amount_of_employees, registered, type)
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.DB.Exec(query, company.ID, company.Name, company.Description, company.AmountEmployees, company.Registered, company.Type)
	return err
}

func (r *CompanyRepository) Get(id string) (*models.Company, error) {
	company := models.Company{}
	query := `SELECT id, name, description, amount_of_employees, registered, type FROM companies WHERE id=$1`
	err := r.DB.QueryRow(query, id).Scan(&company.ID, &company.Name, &company.Description, &company.AmountEmployees, &company.Registered, &company.Type)

	if err == sql.ErrNoRows {
		return nil, errors.New("company not found")
	}
	return &company, err
}

func (r *CompanyRepository) Delete(id string) error {
	query := `DELETE FROM companies WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *CompanyRepository) Update(company *models.Company) error {
	query := `UPDATE companies SET name=$1, description=$2, amount_of_employees=$3, registered=$4, type=$5 WHERE id=$6`
	_, err := r.DB.Exec(query, company.Name, company.Description, company.AmountEmployees, company.Registered, company.Type, company.ID)
	return err
}
