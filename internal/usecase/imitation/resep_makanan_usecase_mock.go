package imitation

import "fits-health/internal/entity"

type RepoResepMockTraditional struct {
	Create func(resep entity.ResepMakanan) error
	GetAll func() []entity.ResepMakanan
	GetID  func(id int) (entity.ResepMakanan, error)
	Update func(id int, resep entity.ResepMakanan) error
	Delete func(id int) error
}

func (r *RepoResepMockTraditional) CreateResep(resep entity.ResepMakanan) error {
	return r.Create(resep)
}

func (r *RepoResepMockTraditional) GetAllResep() []entity.ResepMakanan {
	return r.GetAll()
}

func (r *RepoResepMockTraditional) GetResepByID(id int) (resep entity.ResepMakanan, err error) {
	return r.GetID(id)
}

func (r *RepoResepMockTraditional) UpdateResepByID(id int, resep entity.ResepMakanan) error {
	return r.Update(id, resep)
}

func (r *RepoResepMockTraditional) DeleteResepByID(id int) error {
	return r.Delete(id)
}
