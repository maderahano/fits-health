package domain

import "fits-health/internal/entity"

type AdapterIMTRepository interface {
	CreateIMT(imt entity.IMT) error
	GetAllIMT() []entity.IMT
	GetIMTByID(id int) (imt entity.IMT, err error)
	UpdateIMTByID(id int, imt entity.IMT) error
	DeleteIMTByID(id int) error
}

type AdapterIMTService interface {
	CreateIMTService(imt entity.IMT) error
	GetAllIMTService() []entity.IMT
	GetIMTByIDService(id int) (entity.IMT, error)
	UpdateIMTByIDService(id int, imt entity.IMT) error
	DeleteIMTByIDService(id int) error
}
