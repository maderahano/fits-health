package imitation

import "fits-health/internal/entity"

type RepoProfilMockTraditional struct {
	Create func(profil entity.ProfilKesehatan) error
	GetAll func() []entity.ProfilKesehatan
	GetID  func(id int) (entity.ProfilKesehatan, error)
	Update func(id int, profil entity.ProfilKesehatan) error
	Delete func(id int) error
}

func (r *RepoProfilMockTraditional) CreateProfilKesehatan(profil entity.ProfilKesehatan) error {
	return r.Create(profil)
}

func (r *RepoProfilMockTraditional) GetAllProfilKesehatan() []entity.ProfilKesehatan {
	return r.GetAll()
}

func (r *RepoProfilMockTraditional) GetProfilKesehatanByID(id int) (entity.ProfilKesehatan, error) {
	return r.GetID(id)
}

func (r *RepoProfilMockTraditional) UpdateProfilKesehatanByID(id int, profil entity.ProfilKesehatan) error {
	return r.Update(id, profil)
}

func (r *RepoProfilMockTraditional) DeleteProfilKesehatanByID(id int) error {
	return r.Delete(id)
}
