package services

import (
	"padepokan/app/dto"
	"padepokan/app/models"
	"padepokan/app/repositories"
)

type INasabahService interface {
	GetNasabah() ([]dto.ResGetNasabah, error)
	GetPointNasabah() ([]dto.ResGetPointNasabah, error)
	CreateNasabah(data dto.Nasabah) (models.Nasabah, error)
}

type nasabahService struct {
	nasabahRepo repositories.INasabahRepository
}

func NewNasabahService(nasabahRepo repositories.INasabahRepository) *nasabahService {
	return &nasabahService{nasabahRepo: nasabahRepo}
}

func (c *nasabahService) GetNasabah() ([]dto.ResGetNasabah, error) {
	return c.nasabahRepo.GetNasabah()
}
func (c *nasabahService) GetPointNasabah() ([]dto.ResGetPointNasabah, error) {
	return c.nasabahRepo.GetPointNasabah()
}

func (c *nasabahService) CreateNasabah(data dto.Nasabah) (models.Nasabah, error) {
	dataNasabah := models.Nasabah{Name: data.Name}
	return c.nasabahRepo.CreateNasabah(dataNasabah)
}
