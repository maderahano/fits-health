package usecase

import (
	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type serviceOlahraga struct {
	c    config.Config
	repo domain.AdapterOlahragaRepository
}

func (s *serviceOlahraga) CreateOlahragaService(olahraga entity.Olahraga) error {
	return s.repo.CreateOlahraga(olahraga)
}

func (s *serviceOlahraga) GetAllOlahragaService() []entity.Olahraga {
	return s.repo.GetAllOlahraga()
}

func (s *serviceOlahraga) GetOlahragaByIDService(id int) (entity.Olahraga, error) {
	return s.repo.GetOlahragaByID(id)
}

func (s *serviceOlahraga) UpdateOlahragaByIDService(id int, olahraga entity.Olahraga) error {
	return s.repo.UpdateOlahragaByID(id, olahraga)
}

func (s *serviceOlahraga) DeleteOlahragaByIDService(id int) error {
	return s.repo.DeleteOlahragaByID(id)
}

func NewServiceOlahraga(repo domain.AdapterOlahragaRepository, c config.Config) domain.AdapterOlahragaService {
	return &serviceOlahraga{
		repo: repo,
		c:    c,
	}
}
