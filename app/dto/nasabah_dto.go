package dto

type Nasabah struct {
	Name string `json:"name" validate:"required"`
}

type ResGetNasabah struct {
	AccountId uint64 `json:"account_id" `
	Name      string `json:"name"`
}

type ResGetPointNasabah struct {
	AccountId  uint64 `json:"account_id" `
	Name       string `json:"name"`
	TotalPoint uint32 `json:"total_point"`
}
