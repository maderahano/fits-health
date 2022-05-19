package usecase

import (
	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type serviceJadwal struct {
	c    config.Config
	repo domain.AdapterJadwalRepository
}

func (s *serviceJadwal) CreateJadwalService(jadwal entity.JadwalMakanan) error {
	return s.repo.CreateJadwal(jadwal)
}

func (s *serviceJadwal) GetAllJadwalService() []entity.JadwalMakanan {
	return s.repo.GetAllJadwal()
}

func (s *serviceJadwal) GetJadwalByIDService(id int) (entity.JadwalMakanan, error) {
	return s.repo.GetJadwalByID(id)
}

func (s *serviceJadwal) UpdateJadwalByIDService(id int, jadwal entity.JadwalMakanan) error {
	return s.repo.UpdateJadwalByID(id, jadwal)
}

func (s *serviceJadwal) DeleteJadwalByIDService(id int) error {
	return s.repo.DeleteJadwalByID(id)
}

func NewServiceJadwal(repo domain.AdapterJadwalRepository, c config.Config) domain.AdapterJadwalService {
	return &serviceJadwal{
		repo: repo,
		c:    c,
	}
}
