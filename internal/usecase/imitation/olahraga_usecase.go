package imitation

import "fits-health/internal/entity"

type RepoOlahragaMockTraditional struct {
	Create func(olahraga entity.Olahraga) error
	GetAll func() []entity.Olahraga
	GetID  func(id int) (entity.Olahraga, error)
	Update func(id int, olahraga entity.Olahraga) error
	Delete func(id int) error
}

func (r *RepoOlahragaMockTraditional) CreateOlahraga(olahraga entity.Olahraga) error {
	return r.Create(olahraga)
}

func (r *RepoOlahragaMockTraditional) GetAllOlahraga() []entity.Olahraga {
	return r.GetAll()
}

func (r *RepoOlahragaMockTraditional) GetOlahragaByID(id int) (olahraga entity.Olahraga, err error) {
	return r.GetID(id)
}

func (r *RepoOlahragaMockTraditional) UpdateOlahragaByID(id int, olahraga entity.Olahraga) error {
	return r.Update(id, olahraga)
}

func (r *RepoOlahragaMockTraditional) DeleteOlahragaByID(id int) error {
	return r.Delete(id)
}
