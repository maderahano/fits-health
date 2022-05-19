package repository

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateResep(resep entity.ResepMakanan) error {
	res := r.DB.Create(&resep)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllResep() []entity.ResepMakanan {
	reseps := []entity.ResepMakanan{}
	r.DB.Find(&reseps)

	return reseps
}

func (r *RepositoryMysqlLayer) GetResepByID(id int) (resep entity.ResepMakanan, err error) {
	res := r.DB.Where("id = ?", id).Find(&resep)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateResepByID(id int, resep entity.ResepMakanan) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&resep)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteResepByID(id int) error {
	res := r.DB.Delete(&entity.ResepMakanan{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlResepRepository(db *gorm.DB) domain.AdapterResepRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
