package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockOlahragaService struct {
	mock.Mock
}

func (r *MockOlahragaService) CreateOlahragaService(olahraga entity.Olahraga) error {
	ret := r.Called(olahraga)

	var err error
	if res, ok := ret.Get(0).(func(entity.Olahraga) error); ok {
		err = res(olahraga)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockOlahragaService) GetAllOlahragaService() []entity.Olahraga {
	ret := r.Called()

	var arr []entity.Olahraga
	if res, ok := ret.Get(0).(func() []entity.Olahraga); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.Olahraga)
		}
	}

	return arr
}

func (r *MockOlahragaService) GetOlahragaByIDService(id int) (entity.Olahraga, error) {
	ret := r.Called(id)

	var olahraga entity.Olahraga
	if res, ok := ret.Get(0).(func(int) entity.Olahraga); ok {
		olahraga = res(id)
	} else {
		olahraga = ret.Get(0).(entity.Olahraga)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return olahraga, err
}

func (r *MockOlahragaService) UpdateOlahragaByIDService(id int, olahraga entity.Olahraga) error {
	ret := r.Called(id, olahraga)

	var err error
	if res, ok := ret.Get(0).(func(int, entity.Olahraga) error); ok {
		err = res(id, olahraga)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockOlahragaService) DeleteOlahragaByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
