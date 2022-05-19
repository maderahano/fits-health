package usecase

import (
	"errors"
	"net/http"

	"fits-health/config"
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"fits-health/internal/helper"
)

type serviceUser struct {
	c    config.Config
	repo domain.AdapterUserRepository
}

func (s *serviceUser) RegisterUserService(user entity.User) error {
	return s.repo.RegisterUser(user)
}

func (s *serviceUser) LoginUserService(email, password string) (string, int) {
	user, _ := s.repo.GetUserByEmail(email)
	if (user.Password != password) || (user == entity.User{}) {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(user.ID, user.Email, s.c.JWT_KEY)
	if err != nil {
		return "", http.StatusInternalServerError
	}

	return token, http.StatusOK
}

func (s *serviceUser) GetAllUsersService() []entity.User {
	return s.repo.GetAllUsers()
}

func (s *serviceUser) GetUserByIDService(id int) (entity.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *serviceUser) UpdateUserByIDService(id, idToken int, user entity.User) error {
	if id != idToken {
		return errors.New("error")
	}

	return s.repo.UpdateUserByID(id, user)
}

func (s *serviceUser) DeleteUserByIDService(id int) error {
	return s.repo.DeleteUserByID(id)
}

func NewServiceUser(repo domain.AdapterUserRepository, c config.Config) domain.AdapterUserService {
	return &serviceUser{
		repo: repo,
		c:    c,
	}
}
