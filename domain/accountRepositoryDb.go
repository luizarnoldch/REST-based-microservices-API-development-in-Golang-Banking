package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking-Lib/errs"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang-Banking-Lib/logger"
)

type AccountRespositoryDb struct {
	client *sqlx.DB
}

func (d AccountRespositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts SET customer_id=?, opening_date=?, account_type=?, amount=?,status=?"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inset id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRespositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	result, _ := tx.Exec("INSERT INTO transactions SET account_id=?, amount=?, transaction_type=?, transaction_date=?", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id=?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id=?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRespositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT * FROM accounts WHERE account_id=?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRespositoryDb {
	return AccountRespositoryDb{dbClient}
}
