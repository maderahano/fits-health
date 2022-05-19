package repository

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateIMT(imt entity.IMT) error {
	res := r.DB.Create(&imt)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllIMT() []entity.IMT {
	imts := []entity.IMT{}
	r.DB.Find(&imts)

	return imts
}

func (r *RepositoryMysqlLayer) GetIMTByID(id int) (imt entity.IMT, err error) {
	res := r.DB.Where("id = ?", id).Find(&imt)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateIMTByID(id int, imt entity.IMT) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&imt)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteIMTByID(id int) error {
	res := r.DB.Delete(&entity.IMT{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlIMTRepository(db *gorm.DB) domain.AdapterIMTRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
