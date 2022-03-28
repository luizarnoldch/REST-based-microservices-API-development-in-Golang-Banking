package dto

import "github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking-lib/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transactionDate"`
	CustomerId      string  `json:"-"`
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t TransactionRequest) IsTransactionTypeDeposit() bool {
	return t.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.IsTransactionTypeWithdrawal() && r.IsTransactionTypeDeposit() {
		return errs.NewValidationError("Transaction type can only be deposti or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transactionDate"`
}
