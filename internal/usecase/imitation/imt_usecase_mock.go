package imitation

import (
	"fits-health/internal/entity"
)

type RepoIMTMockTraditional struct {
	Create func(imt entity.IMT) error
	GetAll func() []entity.IMT
	GetID  func(id int) (entity.IMT, error)
	Update func(id int, imt entity.IMT) error
	Delete func(id int) error
}

func (r *RepoIMTMockTraditional) CreateIMT(imt entity.IMT) error {
	return r.Create(imt)
}

func (r *RepoIMTMockTraditional) GetAllIMT() []entity.IMT {
	return r.GetAll()
}

func (r *RepoIMTMockTraditional) GetIMTByID(id int) (imt entity.IMT, err error) {
	return r.GetID(id)
}

func (r *RepoIMTMockTraditional) UpdateIMTByID(id int, imt entity.IMT) error {
	return r.Update(id, imt)
}

func (r *RepoIMTMockTraditional) DeleteIMTByID(id int) error {
	return r.Delete(id)
}
