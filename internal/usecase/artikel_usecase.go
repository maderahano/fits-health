package usecase

import (
	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type serviceArtikel struct {
	c    config.Config
	repo domain.AdapterArtikelRepository
}

func (s *serviceArtikel) CreateArtikelService(artikel entity.Artikel) error {
	return s.repo.CreateArtikel(artikel)
}

func (s *serviceArtikel) GetAllArtikelService() []entity.Artikel {
	return s.repo.GetAllArtikel()
}

func (s *serviceArtikel) GetArtikelByIDService(id int) (entity.Artikel, error) {
	return s.repo.GetArtikelByID(id)
}

func (s *serviceArtikel) UpdateArtikelByIDService(id int, artikel entity.Artikel) error {
	return s.repo.UpdateArtikelByID(id, artikel)
}

func (s *serviceArtikel) DeleteArtikelByIDService(id int) error {
	return s.repo.DeleteArtikelByID(id)
}

func NewServiceArtikel(repo domain.AdapterArtikelRepository, c config.Config) domain.AdapterArtikelService {
	return &serviceArtikel{
		repo: repo,
		c:    c,
	}
}
