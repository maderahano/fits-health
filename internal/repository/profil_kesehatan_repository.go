package repository

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateProfilKesehatan(profil entity.ProfilKesehatan) error {
	res := r.DB.Create(&profil)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllProfilKesehatan() []entity.ProfilKesehatan {
	profils := []entity.ProfilKesehatan{}
	r.DB.Find(&profils)

	return profils
}

func (r *RepositoryMysqlLayer) GetProfilKesehatanByID(id int) (profil entity.ProfilKesehatan, err error) {
	res := r.DB.Where("id = ?", id).Find(&profil)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateProfilKesehatanByID(id int, profil entity.ProfilKesehatan) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&profil)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteProfilKesehatanByID(id int) error {
	res := r.DB.Delete(&entity.ProfilKesehatan{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlProfilKesehatanRepository(db *gorm.DB) domain.AdapterProfilKesehatanRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
