package models

import "time"

type Transaction struct {
	Id                uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountId         uint64    `json:"account_id" gorm:"not null"`
	Description       string    `json:"description" gorm:"type:text"`
	DebitCreditStatus string    `json:"debit_credit_status" gorm:"not null"`
	Amount            uint64    `json:"amount" gorm:"not null"`
	TransactionDate   time.Time `json:"transaction_date" gorm:"not null"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
