package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockResepService struct {
	mock.Mock
}

func (r *MockResepService) CreateResepService(resep entity.ResepMakanan) error {
	ret := r.Called(resep)

	var err error
	if res, ok := ret.Get(0).(func(entity.ResepMakanan) error); ok {
		err = res(resep)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockResepService) GetAllResepService() []entity.ResepMakanan {
	ret := r.Called()

	var arr []entity.ResepMakanan
	if res, ok := ret.Get(0).(func() []entity.ResepMakanan); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.ResepMakanan)
		}
	}

	return arr
}

func (r *MockResepService) GetResepByIDService(id int) (entity.ResepMakanan, error) {
	ret := r.Called(id)

	var resep entity.ResepMakanan
	if res, ok := ret.Get(0).(func(int) entity.ResepMakanan); ok {
		resep = res(id)
	} else {
		resep = ret.Get(0).(entity.ResepMakanan)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return resep, err
}

func (r *MockResepService) UpdateResepByIDService(id int, resep entity.ResepMakanan) error {
	ret := r.Called(id, resep)

	var err error
	if res, ok := ret.Get(0).(func(int, entity.ResepMakanan) error); ok {
		err = res(id, resep)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockResepService) DeleteResepByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
