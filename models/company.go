package models

type Company struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description,omitempty"`
	AmountEmployees int    `json:"amount_of_employees"`
	Registered      bool   `json:"registered"`
	Type            string `json:"type"`
}
