package usecase

import (
	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type serviceIMT struct {
	c    config.Config
	repo domain.AdapterIMTRepository
}

func (s *serviceIMT) CreateIMTService(imt entity.IMT) error {
	return s.repo.CreateIMT(imt)
}

func (s *serviceIMT) GetAllIMTService() []entity.IMT {
	return s.repo.GetAllIMT()
}

func (s *serviceIMT) GetIMTByIDService(id int) (entity.IMT, error) {
	return s.repo.GetIMTByID(id)
}

func (s *serviceIMT) UpdateIMTByIDService(id int, imt entity.IMT) error {
	return s.repo.UpdateIMTByID(id, imt)
}

func (s *serviceIMT) DeleteIMTByIDService(id int) error {
	return s.repo.DeleteIMTByID(id)
}

func NewServiceIMT(repo domain.AdapterIMTRepository, c config.Config) domain.AdapterIMTService {
	return &serviceIMT{
		repo: repo,
		c:    c,
	}
}
