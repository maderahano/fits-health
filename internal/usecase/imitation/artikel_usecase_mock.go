package imitation

import "fits-health/internal/entity"

type RepoArtikelMockTraditional struct {
	Create func(artikel entity.Artikel) error
	GetAll func() []entity.Artikel
	GetID  func(id int) (entity.Artikel, error)
	Update func(id int, artikel entity.Artikel) error
	Delete func(id int) error
}

func (r *RepoArtikelMockTraditional) CreateArtikel(artikel entity.Artikel) error {
	return r.Create(artikel)
}

func (r *RepoArtikelMockTraditional) GetAllArtikel() []entity.Artikel {
	return r.GetAll()
}

func (r *RepoArtikelMockTraditional) GetArtikelByID(id int) (artikel entity.Artikel, err error) {
	return r.GetID(id)
}

func (r *RepoArtikelMockTraditional) UpdateArtikelByID(id int, artikel entity.Artikel) error {
	return r.Update(id, artikel)
}

func (r *RepoArtikelMockTraditional) DeleteArtikelByID(id int) error {
	return r.Delete(id)
}
