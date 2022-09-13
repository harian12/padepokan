package services

import (
	"padepokan/app/dto"
	"padepokan/app/models"
	"padepokan/app/repositories"
	"padepokan/logger"
	"time"
)

type ITransactionService interface {
	CreateTransaction(dataTrx dto.Transaction) (models.Transaction, error)
}

type transactionService struct {
	transactionRepo repositories.ITransactionRepository
}

func NewTransactionService(transactionRepo repositories.ITransactionRepository) *transactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s *transactionService) CreateTransaction(dataTrx dto.Transaction) (models.Transaction, error) {

	dateTransaction, err := time.Parse("2006-01-02", dataTrx.TransactionDate)

	if err != nil {
		logger.ErrorFunc("time parse "+err.Error(), err)
		return models.Transaction{}, err
	}

	transaction := models.Transaction{
		AccountId:         dataTrx.AccountId,
		Description:       dataTrx.Description,
		DebitCreditStatus: dataTrx.DebitCreditStatus,
		Amount:            dataTrx.Amount,
		TransactionDate:   dateTransaction,
	}

	var totalPoin uint32
	amount := dataTrx.Amount
	switch dataTrx.Description {
	case "Beli Pulsa":
		if dataTrx.Amount <= 10000 {
			totalPoin = 0
		} else if dataTrx.Amount > 10000 && dataTrx.Amount <= 30000 {
			amount -= 10000
			totalPoin += (10000 / 1000) * 0
			totalPoin += uint32(float64((amount / 1000)) * 1)
		} else if dataTrx.Amount > 30000 {
			amount -= 10000
			totalPoin += (10000 / 1000) * 0
			amount -= 10000
			totalPoin += (10000 / 1000) * 1
			totalPoin += uint32((float64(amount) / float64(1000)) * 2)
		}
	case "Bayar Listrik":
		if dataTrx.Amount <= 50000 {
			totalPoin = 0
		} else if dataTrx.Amount > 50000 && dataTrx.Amount <= 100000 {
			amount -= 50000
			totalPoin += (50000 / 2000) * 0
			totalPoin += uint32(float64((amount / 2000)) * 1)
		} else if dataTrx.Amount > 100000 {
			amount -= 50000
			totalPoin += (50000 / 2000) * 0
			amount -= 50000
			totalPoin += (50000 / 2000) * 1
			totalPoin += uint32((float64(amount) / float64(2000)) * 2)
		}
	default:
		totalPoin = 0
	}

	data, err := s.transactionRepo.CreateTransaction(transaction, totalPoin)
	return data, err
}
