package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	DateofBirth string `json:"date_of_birth"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	Status      string `json:"status"`
}