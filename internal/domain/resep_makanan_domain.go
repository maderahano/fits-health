package domain

import "fits-health/internal/entity"

type AdapterResepRepository interface {
	CreateResep(resep entity.ResepMakanan) error
	GetAllResep() []entity.ResepMakanan
	GetResepByID(id int) (resep entity.ResepMakanan, err error)
	UpdateResepByID(id int, resep entity.ResepMakanan) error
	DeleteResepByID(id int) error
}

type AdapterResepService interface {
	CreateResepService(resep entity.ResepMakanan) error
	GetAllResepService() []entity.ResepMakanan
	GetResepByIDService(id int) (entity.ResepMakanan, error)
	UpdateResepByIDService(id int, resep entity.ResepMakanan) error
	DeleteResepByIDService(id int) error
}
