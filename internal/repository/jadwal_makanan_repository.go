package repository

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateJadwal(jadwal entity.JadwalMakanan) error {
	res := r.DB.Create(&jadwal)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllJadwal() []entity.JadwalMakanan {
	jadwals := []entity.JadwalMakanan{}
	r.DB.Find(&jadwals)

	return jadwals
}

func (r *RepositoryMysqlLayer) GetJadwalByID(id int) (jadwal entity.JadwalMakanan, err error) {
	res := r.DB.Where("id = ?", id).Find(&jadwal)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateJadwalByID(id int, jadwal entity.JadwalMakanan) error {
	res := r.DB.Where("id = ?").UpdateColumns(&jadwal)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteJadwalByID(id int) error {
	res := r.DB.Delete(&entity.JadwalMakanan{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlJadwalRepository(db *gorm.DB) domain.AdapterJadwalRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
