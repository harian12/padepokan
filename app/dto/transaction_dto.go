package dto

type Transaction struct {
	AccountId         uint64 `json:"account_id" validate:"required,number"`
	TransactionDate   string `json:"transaction_date" validate:"required"`
	Description       string `json:"description" validate:"required"`
	DebitCreditStatus string `json:"debit_credit_status" validate:"required"`
	Amount            uint64 `json:"amount" validate:"required,number"`
}

type Report struct {
	AccountId uint64 `json:"account_id" validate:"required,number"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type ResReport struct {
	TransactionDate string `json:"transaction_date"`
	Description     string `json:"description"`
	Amount          string `json:"amount"`
	Debit           string `json:"debit"`
	Credit          string `json:"credit"`
}
