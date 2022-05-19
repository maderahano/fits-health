package usecase

import (
	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type serviceResep struct {
	c    config.Config
	repo domain.AdapterResepRepository
}

func (s *serviceResep) CreateResepService(resep entity.ResepMakanan) error {
	return s.repo.CreateResep(resep)
}

func (s *serviceResep) GetAllResepService() []entity.ResepMakanan {
	return s.repo.GetAllResep()
}

func (s *serviceResep) GetResepByIDService(id int) (entity.ResepMakanan, error) {
	return s.repo.GetResepByID(id)
}

func (s *serviceResep) UpdateResepByIDService(id int, resep entity.ResepMakanan) error {
	return s.repo.UpdateResepByID(id, resep)
}

func (s *serviceResep) DeleteResepByIDService(id int) error {
	return s.repo.DeleteResepByID(id)
}

func NewServiceResep(repo domain.AdapterResepRepository, c config.Config) domain.AdapterResepService {
	return &serviceResep{
		repo: repo,
		c:    c,
	}
}
