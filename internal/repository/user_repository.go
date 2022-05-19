package repository

import (
	"fmt"

	"gorm.io/gorm"

	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type RepositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *RepositoryMysqlLayer) RegisterUser(user entity.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllUsers() []entity.User {
	users := []entity.User{}
	r.DB.Find(&users)

	return users
}

func (r *RepositoryMysqlLayer) GetUserByID(id int) (user entity.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetUserByEmail(email string) (user entity.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("email not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateUserByID(id int, user entity.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteUserByID(id int) error {
	res := r.DB.Delete(&entity.User{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlUserRepository(db *gorm.DB) domain.AdapterUserRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
