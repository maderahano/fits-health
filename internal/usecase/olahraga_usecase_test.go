package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOlahragaService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(olahraga entity.Olahraga) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(olahraga entity.Olahraga) error { return nil },
			noError: true,
		},
		{
			name: "error internal",
			f: func(olahraga entity.Olahraga) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := imitation.RepoOlahragaMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Create = v.f

			svc := NewServiceOlahraga(&repo, config.Config{})
			err := svc.CreateOlahragaService(entity.Olahraga{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllOlahragaService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.Olahraga
	}{
		{
			name: "success",
			f:    func() []entity.Olahraga { return []entity.Olahraga{} },
		},
		{
			name: "error internal",
			f:    func() []entity.Olahraga { return []entity.Olahraga{} },
		},
	}
	repo := imitation.RepoOlahragaMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceOlahraga(&repo, config.Config{})
			svc.GetAllOlahragaService()
		})
	}
}

func TestGetOlahragaByIDService(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.Olahraga, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.Olahraga, error) { return entity.Olahraga{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.Olahraga, error) { return entity.Olahraga{}, errors.New("error") },
		},
	}
	repo := imitation.RepoOlahragaMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceOlahraga(&repo, config.Config{})
			_, err := svc.GetOlahragaByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateOlahragaByIDService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(id int, olahraga entity.Olahraga) error
		noError bool
		id      int
	}{
		{
			name:    "success",
			f:       func(id int, olahraga entity.Olahraga) error { return nil },
			noError: true,
			id:      1,
		},
		{
			name: "error internal",
			f: func(id int, olahraga entity.Olahraga) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repo := imitation.RepoOlahragaMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceOlahraga(&repo, config.Config{})
			err := svc.UpdateOlahragaByIDService(v.id, entity.Olahraga{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}

}

func TestDeleteOlahragaByIDService(t *testing.T) {
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
	repo := imitation.RepoOlahragaMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceOlahraga(&repo, config.Config{})
			err := svc.DeleteOlahragaByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
