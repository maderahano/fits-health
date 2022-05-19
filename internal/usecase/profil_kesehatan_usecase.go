package usecase

import (
	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type serviceProfilKesehatan struct {
	c    config.Config
	repo domain.AdapterProfilKesehatanRepository
}

func (s *serviceProfilKesehatan) CreateProfilKesehatanService(profil entity.ProfilKesehatan) error {
	return s.repo.CreateProfilKesehatan(profil)
}

func (s *serviceProfilKesehatan) GetAllProfilKesehatanService() []entity.ProfilKesehatan {
	return s.repo.GetAllProfilKesehatan()
}

func (s *serviceProfilKesehatan) GetProfilKesehatanByIDService(id int) (entity.ProfilKesehatan, error) {
	return s.repo.GetProfilKesehatanByID(id)
}

func (s *serviceProfilKesehatan) UpdateProfilKesehatanByIDService(id int, profil entity.ProfilKesehatan) error {
	return s.repo.UpdateProfilKesehatanByID(id, profil)
}

func (s *serviceProfilKesehatan) DeleteProfilKesehatanByIDService(id int) error {
	return s.repo.DeleteProfilKesehatanByID(id)
}

func NewServiceProfileKesehatan(repo domain.AdapterProfilKesehatanRepository, c config.Config) domain.AdapterProfilKesehatanService {
	return &serviceProfilKesehatan{
		repo: repo,
		c:    c,
	}
}
