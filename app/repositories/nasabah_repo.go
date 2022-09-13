package repositories

import (
	"gorm.io/gorm"
	"padepokan/app/models"
)

type INasabahRepository interface {
	GetNasabah() ([]models.Nasabah, error)
	CreateNasabah(data models.Nasabah) (models.Nasabah, error)
}

type nasabahRepository struct {
	db *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) *nasabahRepository {
	return &nasabahRepository{db}
}

func (r *nasabahRepository) GetNasabah() ([]models.Nasabah, error) {
	var nasabah []models.Nasabah
	err := r.db.Find(&nasabah).Error
	return nasabah, err
}

func (r *nasabahRepository) CreateNasabah(data models.Nasabah) (models.Nasabah, error) {
	err := r.db.Create(&data).Error
	return data, err
}
