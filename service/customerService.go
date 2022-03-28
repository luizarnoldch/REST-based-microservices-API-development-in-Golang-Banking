package service

import (
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking-Lib/errs"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking/domain"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking/dto"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	//GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	//func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
