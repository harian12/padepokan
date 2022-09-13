package services

import (
	"padepokan/app/dto"
	"padepokan/app/models"
	"padepokan/app/repositories"
)

type INasabahService interface {
	GetNasabah() ([]models.Nasabah, error)
	CreateNasabah(data dto.Nasabah) (models.Nasabah, error)
}

type nasabahService struct {
	nasabahRepo repositories.INasabahRepository
}

func NewNasabahService(nasabahRepo repositories.INasabahRepository) *nasabahService {
	return &nasabahService{nasabahRepo: nasabahRepo}
}

func (c *nasabahService) GetNasabah() ([]models.Nasabah, error) {
	return c.nasabahRepo.GetNasabah()
}

func (c *nasabahService) CreateNasabah(data dto.Nasabah) (models.Nasabah, error) {
	dataNasabah := models.Nasabah{Name: data.Name}
	return c.nasabahRepo.CreateNasabah(dataNasabah)
}
