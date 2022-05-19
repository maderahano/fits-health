package imitation

import "fits-health/internal/entity"

type RepoJadwalMockTraditional struct {
	Create func(jadwal entity.JadwalMakanan) error
	GetAll func() []entity.JadwalMakanan
	GetID  func(id int) (entity.JadwalMakanan, error)
	Update func(id int, jadwal entity.JadwalMakanan) error
	Delete func(id int) error
}

func (r *RepoJadwalMockTraditional) CreateJadwal(jadwal entity.JadwalMakanan) error {
	return r.Create(jadwal)
}

func (r *RepoJadwalMockTraditional) GetAllJadwal() []entity.JadwalMakanan {
	return r.GetAll()
}

func (r *RepoJadwalMockTraditional) GetJadwalByID(id int) (jadwal entity.JadwalMakanan, err error) {
	return r.GetID(id)
}

func (r *RepoJadwalMockTraditional) UpdateJadwalByID(id int, jadwal entity.JadwalMakanan) error {
	return r.Update(id, jadwal)
}

func (r *RepoJadwalMockTraditional) DeleteJadwalByID(id int) error {
	return r.Delete(id)
}
