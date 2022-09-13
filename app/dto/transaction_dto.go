package dto

type Transaction struct {
	AccountId         uint64 `json:"account_id" validate:"required,number"`
	TransactionDate   string `json:"transaction_date" validate:"required"`
	Description       string `json:"description" validate:"required"`
	DebitCreditStatus string `json:"debit_credit_status" validate:"required"`
	Amount            uint64 `json:"amount" validate:"required,number"`
}
