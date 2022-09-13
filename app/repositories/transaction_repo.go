package repositories

import (
	"gorm.io/gorm"
	"padepokan/app/models"
)

type ITransactionRepository interface {
	CreateTransaction(transaction models.Transaction, point uint32) (models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction models.Transaction, point uint32) (models.Transaction, error) {

	trx := r.db.Begin()
	err := trx.Create(&transaction).Error
	if err != nil {
		trx.Rollback()
		return models.Transaction{}, err
	}

	var nasabah models.Nasabah
	err = r.db.First(&nasabah, transaction.AccountId).Error
	if err != nil {
		trx.Rollback()
		return models.Transaction{}, err
	}

	nasabah.TotalPoint = nasabah.TotalPoint + point

	err = trx.Save(&nasabah).Error
	if err != nil {
		trx.Rollback()
		return models.Transaction{}, err
	}
	err = trx.Commit().Error
	return transaction, err
}
