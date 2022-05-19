package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProfilKesehatanService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(profil entity.ProfilKesehatan) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(profil entity.ProfilKesehatan) error { return nil },
			noError: true,
		},
		{
			name: "error internal",
			f: func(profil entity.ProfilKesehatan) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := imitation.RepoProfilMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Create = v.f

			svc := NewServiceProfileKesehatan(&repo, config.Config{})
			err := svc.CreateProfilKesehatanService(entity.ProfilKesehatan{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllProfilKesehatanService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.ProfilKesehatan
	}{
		{
			name: "success",
			f:    func() []entity.ProfilKesehatan { return []entity.ProfilKesehatan{} },
		},
		{
			name: "error internal",
			f:    func() []entity.ProfilKesehatan { return []entity.ProfilKesehatan{} },
		},
	}
	repo := imitation.RepoProfilMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceProfileKesehatan(&repo, config.Config{})
			svc.GetAllProfilKesehatanService()
		})
	}
}

func TestGetProfilKesehatanByID(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.ProfilKesehatan, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.ProfilKesehatan, error) { return entity.ProfilKesehatan{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.ProfilKesehatan, error) { return entity.ProfilKesehatan{}, errors.New("error") },
		},
	}
	repo := imitation.RepoProfilMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceProfileKesehatan(&repo, config.Config{})
			_, err := svc.GetProfilKesehatanByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateProfilKesehatanByIDService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(id int, profil entity.ProfilKesehatan) error
		noError bool
		id      int
	}{
		{
			name:    "success",
			f:       func(id int, profil entity.ProfilKesehatan) error { return nil },
			noError: true,
			id:      1,
		},
		{
			name: "error internal",
			f: func(id int, profil entity.ProfilKesehatan) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repo := imitation.RepoProfilMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceProfileKesehatan(&repo, config.Config{})
			err := svc.UpdateProfilKesehatanByIDService(v.id, entity.ProfilKesehatan{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteProfilKesehatanByIDService(t *testing.T) {
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
	repo := imitation.RepoProfilMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceProfileKesehatan(&repo, config.Config{})
			err := svc.DeleteProfilKesehatanByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
