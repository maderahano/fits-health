package domain

import "fits-health/internal/entity"

type AdapterProfilKesehatanRepository interface {
	CreateProfilKesehatan(imt entity.ProfilKesehatan) error
	GetAllProfilKesehatan() []entity.ProfilKesehatan
	GetProfilKesehatanByID(id int) (imt entity.ProfilKesehatan, err error)
	UpdateProfilKesehatanByID(id int, profil entity.ProfilKesehatan) error
	DeleteProfilKesehatanByID(id int) error
}

type AdapterProfilKesehatanService interface {
	CreateProfilKesehatanService(imt entity.ProfilKesehatan) error
	GetAllProfilKesehatanService() []entity.ProfilKesehatan
	GetProfilKesehatanByIDService(id int) (entity.ProfilKesehatan, error)
	UpdateProfilKesehatanByIDService(id int, profil entity.ProfilKesehatan) error
	DeleteProfilKesehatanByIDService(id int) error
}
