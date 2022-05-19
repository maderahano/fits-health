package domain

import "fits-health/internal/entity"

type AdapterUserRepository interface {
	RegisterUser(user entity.User) error
	GetAllUsers() []entity.User
	GetUserByID(id int) (user entity.User, err error)
	GetUserByEmail(email string) (user entity.User, err error)
	UpdateUserByID(id int, user entity.User) error
	DeleteUserByID(id int) error
}

type AdapterUserService interface {
	RegisterUserService(user entity.User) error
	LoginUserService(email, password string) (string, int)
	GetAllUsersService() []entity.User
	GetUserByIDService(id int) (entity.User, error)
	UpdateUserByIDService(id, idToken int, user entity.User) error
	DeleteUserByIDService(id int) error
}
