package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUserService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(user entity.User) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(user entity.User) error { return nil },
			noError: true,
		},
		{
			name: "error internal",
			f: func(user entity.User) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := imitation.RepoUserMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Register = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.RegisterUserService(entity.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestLoginUserService(t *testing.T) {
	testCase := []struct {
		name     string
		email    string
		password string
		f        func(email string) (entity.User, error)
		http     int
	}{
		{
			name:     "success",
			email:    "user@gmail.com",
			password: "000000",
			f:        func(email string) (entity.User, error) { return entity.User{}, nil },
			http:     200,
		},
		{
			name:     "error internal",
			email:    "user@gmail.com",
			password: "000000",
			f: func(email string) (entity.User, error) {
				return entity.User{}, errors.New("error")
			},
			http: 500,
		},
	}
	repo := imitation.RepoUserMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetEmail = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, httpResult := svc.LoginUserService(v.email, v.password)
			if httpResult != 401 {
				t.Errorf("handler returned wrong status code: got %v want %v",
					v.http, httpResult)
			}
		})
	}
}

func TestGetAllUsersService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.User
	}{
		{
			name: "success",
			f:    func() []entity.User { return []entity.User{} },
		},
		{
			name: "error internal",
			f:    func() []entity.User { return []entity.User{} },
		},
	}
	repo := imitation.RepoUserMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceUser(&repo, config.Config{})
			svc.GetAllUsersService()
		})
	}
}

func TestGetUserByIDService(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.User, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.User, error) { return entity.User{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.User, error) { return entity.User{}, errors.New("error") },
		},
	}
	repo := imitation.RepoUserMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceUser(&repo, config.Config{})
			_, err := svc.GetUserByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateUserByIDService(t *testing.T) {
	testCase := []struct {
		name      string
		f         func(id int, user entity.User) error
		noError   bool
		token, id int
	}{
		{
			name:    "success",
			f:       func(id int, user entity.User) error { return nil },
			noError: true,
			token:   1,
			id:      1,
		},
		{
			name:    "error token != id",
			f:       func(id int, user entity.User) error { return nil },
			noError: false,
			token:   1,
			id:      2,
		},
		{
			name: "error internal",
			f: func(id int, user entity.User) error {
				return errors.New("error")
			},
			noError: false,
			token:   1,
			id:      1,
		},
	}
	repo := imitation.RepoUserMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.UpdateUserByIDService(v.id, v.token, entity.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteUserByIDService(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) error
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) error { return nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) error { return errors.New("error") },
		},
	}
	repo := imitation.RepoUserMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceUser(&repo, config.Config{})
			err := svc.DeleteUserByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

// func TestNewServiceUser(t *testing.T) {
// 	testTable := []struct {
// 		name string
// 		f    func(db *gorm.DB) domain.AdapterRepository
// 	}{
// 		{
// 			name: "success",
// 			f:    func(db *gorm.DB) domain.AdapterRepository { return nil },
// 		},
// 	}
// 	repo := mock.RepoMockTraditional{}

// 	for _, v := range testTable {
// 		t.Run(v.name, func(t *testing.T) {
// 			repo.NewRepo = v.f

// 			NewServiceUser(&repo, config.Config{})
// 		})
// 	}
// }
