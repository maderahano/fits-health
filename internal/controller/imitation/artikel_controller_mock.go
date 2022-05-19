package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockArtikelService struct {
	mock.Mock
}

func (r *MockArtikelService) CreateArtikelService(artikel entity.Artikel) error {
	ret := r.Called(artikel)

	var err error
	if res, ok := ret.Get(0).(func(entity.Artikel) error); ok {
		err = res(artikel)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockArtikelService) GetAllArtikelService() []entity.Artikel {
	ret := r.Called()

	var arr []entity.Artikel
	if res, ok := ret.Get(0).(func() []entity.Artikel); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.Artikel)
		}
	}

	return arr
}

func (r *MockArtikelService) GetArtikelByIDService(id int) (entity.Artikel, error) {
	ret := r.Called(id)

	var artikel entity.Artikel
	if res, ok := ret.Get(0).(func(int) entity.Artikel); ok {
		artikel = res(id)
	} else {
		artikel = ret.Get(0).(entity.Artikel)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return artikel, err
}

func (r *MockArtikelService) UpdateArtikelByIDService(id int, artikel entity.Artikel) error {
	ret := r.Called(id, artikel)

	var err error
	if res, ok := ret.Get(0).(func(int, entity.Artikel) error); ok {
		err = res(id, artikel)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockArtikelService) DeleteArtikelByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
