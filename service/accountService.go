package service

import (
	"time"

	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking/domain"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking/dto"
	"github.com/luizarnoldch/REST-based-microservices-API-development-in-Golang/banking-lib/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	
		a := domain.Account{
			AccountId:   "",
			CustomerId:  req.CustomerId,
			OpeningDate: time.Now().Format(dbTSLayout),
			AccountType: req.AccountType,
			Amount:      req.Amount,
			Status:      "1",
		}
	newAccount, err := s.repo.Save(a)
	if  err != nil {
		return nil, err
	} 
	response := newAccount.ToNewAccountResponseDto()
	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appErr := s.repo.SaveTransaction(t)
	if appErr != nil {
		return nil, appErr
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
