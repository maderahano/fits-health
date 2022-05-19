package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArtikelService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(artikel entity.Artikel) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(artikel entity.Artikel) error { return nil },
			noError: true,
		},
		{
			name:    "error internal",
			f:       func(artikel entity.Artikel) error { return errors.New("error") },
			noError: false,
		},
	}
	repo := imitation.RepoArtikelMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Create = v.f

			svc := NewServiceArtikel(&repo, config.Config{})
			err := svc.CreateArtikelService(entity.Artikel{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllArtikelService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.Artikel
	}{
		{
			name: "success",
			f:    func() []entity.Artikel { return []entity.Artikel{} },
		},
		{
			name: "error internal",
			f:    func() []entity.Artikel { return []entity.Artikel{} },
		},
	}
	repo := imitation.RepoArtikelMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceArtikel(&repo, config.Config{})
			svc.GetAllArtikelService()
		})
	}
}

func TestGetArtikelByID(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.Artikel, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.Artikel, error) { return entity.Artikel{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.Artikel, error) { return entity.Artikel{}, errors.New("error") },
		},
	}
	repo := imitation.RepoArtikelMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceArtikel(&repo, config.Config{})
			_, err := svc.GetArtikelByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateArtikelByID(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(id int, artikel entity.Artikel) error
		noError bool
		id      int
	}{
		{
			name:    "success",
			f:       func(id int, artikel entity.Artikel) error { return nil },
			noError: true,
			id:      1,
		},
		{
			name: "error internal",
			f: func(id int, artikel entity.Artikel) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repo := imitation.RepoArtikelMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceArtikel(&repo, config.Config{})
			err := svc.UpdateArtikelByIDService(v.id, entity.Artikel{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteArtikelByIDService(t *testing.T) {
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
	repo := imitation.RepoArtikelMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceArtikel(&repo, config.Config{})
			err := svc.DeleteArtikelByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
