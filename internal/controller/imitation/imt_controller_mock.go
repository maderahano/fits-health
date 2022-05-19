package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockIMTService struct {
	mock.Mock
}

func (r *MockIMTService) CreateIMTService(imt entity.IMT) error {
	ret := r.Called(imt)

	var err error
	if res, ok := ret.Get(0).(func(entity.IMT) error); ok {
		err = res(imt)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockIMTService) GetAllIMTService() []entity.IMT {
	ret := r.Called()

	var arr []entity.IMT
	if res, ok := ret.Get(0).(func() []entity.IMT); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.IMT)
		}
	}

	return arr
}

func (r *MockIMTService) GetIMTByIDService(id int) (entity.IMT, error) {
	ret := r.Called(id)

	var imt entity.IMT
	if res, ok := ret.Get(0).(func(int) entity.IMT); ok {
		imt = res(id)
	} else {
		imt = ret.Get(0).(entity.IMT)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return imt, err
}

func (r *MockIMTService) UpdateIMTByIDService(id int, imt entity.IMT) error {
	ret := r.Called(id, imt)

	var err error
	if res, ok := ret.Get(0).(func(int, entity.IMT) error); ok {
		err = res(id, imt)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockIMTService) DeleteIMTByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
