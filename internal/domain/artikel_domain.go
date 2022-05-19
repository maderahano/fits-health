package domain

import "fits-health/internal/entity"

type AdapterArtikelRepository interface {
	CreateArtikel(artikel entity.Artikel) error
	GetAllArtikel() []entity.Artikel
	GetArtikelByID(id int) (artikel entity.Artikel, err error)
	UpdateArtikelByID(id int, artikel entity.Artikel) error
	DeleteArtikelByID(id int) error
}

type AdapterArtikelService interface {
	CreateArtikelService(artikel entity.Artikel) error
	GetAllArtikelService() []entity.Artikel
	GetArtikelByIDService(id int) (entity.Artikel, error)
	UpdateArtikelByIDService(id int, artikel entity.Artikel) error
	DeleteArtikelByIDService(id int) error
}
