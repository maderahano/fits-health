package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateResepService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(resep entity.ResepMakanan) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(resep entity.ResepMakanan) error { return nil },
			noError: true,
		},
		{
			name: "error internal",
			f: func(resep entity.ResepMakanan) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := imitation.RepoResepMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Create = v.f

			svc := NewServiceResep(&repo, config.Config{})
			err := svc.CreateResepService(entity.ResepMakanan{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllResepService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.ResepMakanan
	}{
		{
			name: "success",
			f:    func() []entity.ResepMakanan { return []entity.ResepMakanan{} },
		},
		{
			name: "error internal",
			f:    func() []entity.ResepMakanan { return []entity.ResepMakanan{} },
		},
	}
	repo := imitation.RepoResepMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceResep(&repo, config.Config{})
			svc.GetAllResepService()
		})
	}
}

func TestGetResepByIDService(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.ResepMakanan, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.ResepMakanan, error) { return entity.ResepMakanan{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.ResepMakanan, error) { return entity.ResepMakanan{}, errors.New("error") },
		},
	}
	repo := imitation.RepoResepMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceResep(&repo, config.Config{})
			_, err := svc.GetResepByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateResepByIDService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(id int, resep entity.ResepMakanan) error
		noError bool
		id      int
	}{
		{
			name:    "success",
			f:       func(id int, resep entity.ResepMakanan) error { return nil },
			noError: true,
			id:      1,
		},
		{
			name: "error internal",
			f: func(id int, resep entity.ResepMakanan) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repo := imitation.RepoResepMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceResep(&repo, config.Config{})
			err := svc.UpdateResepByIDService(v.id, entity.ResepMakanan{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteResepByIDService(t *testing.T) {
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
	repo := imitation.RepoResepMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceResep(&repo, config.Config{})
			err := svc.DeleteResepByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
