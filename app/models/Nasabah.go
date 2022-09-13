package models

import "time"

type Nasabah struct {
	AccountId  uint64    `json:"account_id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"not null;size:50"`
	TotalPoint uint32    `json:"total_point" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
