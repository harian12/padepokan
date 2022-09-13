package dto

type Nasabah struct {
	Name string `json:"name" validate:"required"`
}
