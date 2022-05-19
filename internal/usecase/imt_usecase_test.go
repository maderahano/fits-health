package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateIMTService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(imt entity.IMT) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(imt entity.IMT) error { return nil },
			noError: true,
		},
		{
			name: "error internal",
			f: func(imt entity.IMT) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := imitation.RepoIMTMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Create = v.f

			svc := NewServiceIMT(&repo, config.Config{})
			err := svc.CreateIMTService(entity.IMT{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllIMTService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.IMT
	}{
		{
			name: "success",
			f:    func() []entity.IMT { return []entity.IMT{} },
		},
		{
			name: "error internal",
			f:    func() []entity.IMT { return []entity.IMT{} },
		},
	}
	repo := imitation.RepoIMTMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceIMT(&repo, config.Config{})
			svc.GetAllIMTService()
		})
	}
}

func TestGetIMTByIDService(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.IMT, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.IMT, error) { return entity.IMT{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.IMT, error) { return entity.IMT{}, errors.New("error") },
		},
	}
	repo := imitation.RepoIMTMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceIMT(&repo, config.Config{})
			_, err := svc.GetIMTByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateIMTByIDService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(id int, imt entity.IMT) error
		noError bool
		id      int
	}{
		{
			name:    "success",
			f:       func(id int, imt entity.IMT) error { return nil },
			noError: true,
			id:      1,
		},
		{
			name: "error internal",
			f: func(id int, imt entity.IMT) error {
				return errors.New("error")
			},
			noError: false,
			id:      1,
		},
	}
	repo := imitation.RepoIMTMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceIMT(&repo, config.Config{})
			err := svc.UpdateIMTByIDService(v.id, entity.IMT{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteIMTByIdService(t *testing.T) {
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
			f:       func(id int) error { return errors.New("error internal") },
		},
	}
	repo := imitation.RepoIMTMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceIMT(&repo, config.Config{})
			err := svc.DeleteIMTByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
