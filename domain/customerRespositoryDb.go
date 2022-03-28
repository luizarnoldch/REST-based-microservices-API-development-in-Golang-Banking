package domain

import (
	"database/sql"
	//"fmt"
	//"time"
	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking-lib/errs"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking-lib/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	//findAllSql := "SELECT * FROM customers"
	//rows, err := d.client.Query(findAllSql)
	//var rows *sql.Rows

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select * from customers"
		err = d.client.Select(&customers, findAllSql)
		//rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select * from customers where status = ?"
		//rows, err = d.client.Query(findAllSql, status)
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error While Querying customers" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	/*
		sqlx.StructScan(rows, &customers)
		if err != nil {
			logger.Error("Error while scanning customers customers")
			return nil, errs.NewUnexpectedError("Unexpected database error")
			//return nil, logger.Error("Error while scanning customers")
		}
	*/

	/*
		for rows.Next() {
			var c Customer
			err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
			if err != nil {
				logger.Error("Error while scanning customers customers")
				return nil, errs.NewUnexpectedError("Unexpected database error")
				//return nil, logger.Error("Error while scanning customers")
			}
			customers = append(customers, c)
		}
	*/
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT * FROM customers WHERE customer_id = ?"

	//row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found")
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			//log.Println("Error while scanning customers" + err.Error())
			logger.Error("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
