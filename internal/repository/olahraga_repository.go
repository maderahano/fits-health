package repository

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateOlahraga(olahraga entity.Olahraga) error {
	res := r.DB.Create(&olahraga)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllOlahraga() []entity.Olahraga {
	olahragas := []entity.Olahraga{}
	r.DB.Find(&olahragas)

	return olahragas
}

func (r *RepositoryMysqlLayer) GetOlahragaByID(id int) (olahraga entity.Olahraga, err error) {
	res := r.DB.Where("id = ?", id).Find(&olahraga)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateOlahragaByID(id int, olahraga entity.Olahraga) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&olahraga)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteOlahragaByID(id int) error {
	res := r.DB.Delete(&entity.Olahraga{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlOlahragaRepository(db *gorm.DB) domain.AdapterOlahragaRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
