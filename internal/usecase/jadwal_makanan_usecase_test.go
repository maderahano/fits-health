package usecase

import (
	"errors"
	"fits-health/config"
	"fits-health/internal/entity"
	"fits-health/internal/usecase/imitation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJadwalService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(jadwal entity.JadwalMakanan) error
		noError bool
	}{
		{
			name:    "success",
			f:       func(jadwal entity.JadwalMakanan) error { return nil },
			noError: true,
		},
		{
			name: "error internal",
			f: func(jadwal entity.JadwalMakanan) error {
				return errors.New("error")
			},
			noError: false,
		},
	}
	repo := imitation.RepoJadwalMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Create = v.f

			svc := NewServiceJadwal(&repo, config.Config{})
			err := svc.CreateJadwalService(entity.JadwalMakanan{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAllJadwalService(t *testing.T) {
	testCase := []struct {
		name string
		f    func() []entity.JadwalMakanan
	}{
		{
			name: "success",
			f:    func() []entity.JadwalMakanan { return []entity.JadwalMakanan{} },
		},
		{
			name: "error internal",
			f:    func() []entity.JadwalMakanan { return []entity.JadwalMakanan{} },
		},
	}
	repo := imitation.RepoJadwalMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetAll = v.f

			svc := NewServiceJadwal(&repo, config.Config{})
			svc.GetAllJadwalService()
		})
	}
}

func TestGetJadwalByIDService(t *testing.T) {
	testCase := []struct {
		id      int
		name    string
		noError bool
		f       func(id int) (entity.JadwalMakanan, error)
	}{
		{
			id:      1,
			name:    "success",
			noError: true,
			f:       func(id int) (entity.JadwalMakanan, error) { return entity.JadwalMakanan{}, nil },
		},
		{
			id:      1,
			name:    "error internal",
			noError: false,
			f:       func(id int) (entity.JadwalMakanan, error) { return entity.JadwalMakanan{}, errors.New("error") },
		},
	}
	repo := imitation.RepoJadwalMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.GetID = v.f

			svc := NewServiceJadwal(&repo, config.Config{})
			_, err := svc.GetJadwalByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateJadwalByIDService(t *testing.T) {
	testCase := []struct {
		name    string
		f       func(id int, jadwal entity.JadwalMakanan) error
		noError bool
		id      int
	}{
		{
			name:    "success",
			f:       func(id int, jadwal entity.JadwalMakanan) error { return nil },
			noError: true,
			id:      1,
		},
		{
			name:    "error internal",
			f:       func(id int, jadwal entity.JadwalMakanan) error { return errors.New("error") },
			noError: false,
			id:      1,
		},
	}
	repo := imitation.RepoJadwalMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Update = v.f

			svc := NewServiceJadwal(&repo, config.Config{})
			err := svc.UpdateJadwalByIDService(v.id, entity.JadwalMakanan{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteJadwalByIDService(t *testing.T) {
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
	repo := imitation.RepoJadwalMockTraditional{}

	for _, v := range testCase {
		t.Run(v.name, func(t *testing.T) {
			repo.Delete = v.f

			svc := NewServiceJadwal(&repo, config.Config{})
			err := svc.DeleteJadwalByIDService(v.id)
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
