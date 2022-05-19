package repository

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateArtikel(artikel entity.Artikel) error {
	res := r.DB.Create(&artikel)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllArtikel() []entity.Artikel {
	artikels := []entity.Artikel{}
	r.DB.Find(&artikels)

	return artikels
}

func (r *RepositoryMysqlLayer) GetArtikelByID(id int) (artikel entity.Artikel, err error) {
	res := r.DB.Where("id = ?", id).Find(&artikel)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateArtikelByID(id int, artikel entity.Artikel) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&artikel)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteArtikelByID(id int) error {
	res := r.DB.Delete(&entity.Artikel{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlArtikelRepository(db *gorm.DB) domain.AdapterArtikelRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
