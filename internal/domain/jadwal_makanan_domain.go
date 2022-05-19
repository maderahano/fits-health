package domain

import "fits-health/internal/entity"

type AdapterJadwalRepository interface {
	CreateJadwal(imt entity.JadwalMakanan) error
	GetAllJadwal() []entity.JadwalMakanan
	GetJadwalByID(id int) (jadwal entity.JadwalMakanan, err error)
	UpdateJadwalByID(id int, jadwal entity.JadwalMakanan) error
	DeleteJadwalByID(id int) error
}

type AdapterJadwalService interface {
	CreateJadwalService(jadwal entity.JadwalMakanan) error
	GetAllJadwalService() []entity.JadwalMakanan
	GetJadwalByIDService(id int) (entity.JadwalMakanan, error)
	UpdateJadwalByIDService(id int, jadwal entity.JadwalMakanan) error
	DeleteJadwalByIDService(id int) error
}
