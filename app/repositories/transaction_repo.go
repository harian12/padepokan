package repositories

import (
	"gorm.io/gorm"
	"padepokan/app/dto"
	"padepokan/app/models"
	"padepokan/logger"
	"time"
)

type ITransactionRepository interface {
	CreateTransaction(transaction models.Transaction, point uint32) (models.Transaction, error)
	ReportTransaction(data dto.Report) ([]dto.ResReport, error)
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

func (r *transactionRepository) ReportTransaction(data dto.Report) ([]dto.ResReport, error) {
	start, err := time.Parse("02/01/2006", data.StartDate)
	if err != nil {
		logger.ErrorFunc("time parse Start "+err.Error(), err)
		return nil, err
	}
	end, err := time.Parse("02/01/2006", data.EndDate)
	if err != nil {
		logger.ErrorFunc("time parse End "+err.Error(), err)
		return nil, err
	}
	var trx []dto.ResReport
	err = r.db.Table("transactions").Select("DATE_FORMAT(transaction_date,'%Y-%m-%d') as transaction_date, description, IF(debit_credit_status = 'D',format(amount,0,'id_ID'), '-') as debit,IF(debit_credit_status = 'C',format(amount,0,'id_ID'), '-') as credit, format(amount,0,'id_ID')").Where("account_id = ? and transaction_date between ? and ?", data.AccountId, start, end).Find(&trx).Error
	return trx, err
}
