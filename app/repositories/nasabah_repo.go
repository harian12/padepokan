package repositories

import (
	"gorm.io/gorm"
	"padepokan/app/dto"
	"padepokan/app/models"
)

type INasabahRepository interface {
	GetNasabah() ([]dto.ResGetNasabah, error)
	GetPointNasabah() ([]dto.ResGetPointNasabah, error)
	CreateNasabah(data models.Nasabah) (models.Nasabah, error)
}

type nasabahRepository struct {
	db *gorm.DB
}

func NewNasabahRepository(db *gorm.DB) *nasabahRepository {
	return &nasabahRepository{db}
}

func (r *nasabahRepository) GetNasabah() ([]dto.ResGetNasabah, error) {
	var nasabah []dto.ResGetNasabah
	err := r.db.Table("nasabahs").Select("account_id,name").Find(&nasabah).Error
	return nasabah, err
}
func (r *nasabahRepository) GetPointNasabah() ([]dto.ResGetPointNasabah, error) {
	var nasabah []dto.ResGetPointNasabah
	err := r.db.Table("nasabahs").Select("account_id,name,total_point").Find(&nasabah).Error
	return nasabah, err
}

func (r *nasabahRepository) CreateNasabah(data models.Nasabah) (models.Nasabah, error) {
	err := r.db.Create(&data).Error
	return data, err
}
