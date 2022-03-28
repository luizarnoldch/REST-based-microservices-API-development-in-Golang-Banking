package domain

import (
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking-lib/errs"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking/dto"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	DateofBirth string `db:"date_of_birth"`
	City        string
	Zipcode     string
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	//ById(string) (*Customer, error)
}
