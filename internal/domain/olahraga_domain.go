package domain

import "fits-health/internal/entity"

type AdapterOlahragaRepository interface {
	CreateOlahraga(olahraga entity.Olahraga) error
	GetAllOlahraga() []entity.Olahraga
	GetOlahragaByID(id int) (olahraga entity.Olahraga, err error)
	UpdateOlahragaByID(id int, olahraga entity.Olahraga) error
	DeleteOlahragaByID(id int) error
}

type AdapterOlahragaService interface {
	CreateOlahragaService(olahraga entity.Olahraga) error
	GetAllOlahragaService() []entity.Olahraga
	GetOlahragaByIDService(id int) (entity.Olahraga, error)
	UpdateOlahragaByIDService(id int, olahraga entity.Olahraga) error
	DeleteOlahragaByIDService(id int) error
}
