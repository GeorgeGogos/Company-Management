package services

import (
	"github.com/GeorgeGogos/Company-Management/models"
	"github.com/GeorgeGogos/Company-Management/repository"
)

type CompanyService struct {
	Repo *repository.CompanyRepository
}

func (s *CompanyService) Create(company *models.Company) error {
	return s.Repo.Create(company)
}

func (s *CompanyService) Get(id string) (*models.Company, error) {
	return s.Repo.Get(id)
}

func (s *CompanyService) Update(company *models.Company) error {
	return s.Repo.Update(company)
}

func (s *CompanyService) Delete(id string) error {
	return s.Repo.Delete(id)
}
