package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockJadwalService struct {
	mock.Mock
}

func (r *MockJadwalService) CreateJadwalService(jadwal entity.JadwalMakanan) error {
	ret := r.Called(jadwal)

	var err error
	if res, ok := ret.Get(0).(func(entity.JadwalMakanan) error); ok {
		err = res(jadwal)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockJadwalService) GetAllJadwalService() []entity.JadwalMakanan {
	ret := r.Called()

	var arr []entity.JadwalMakanan
	if res, ok := ret.Get(0).(func() []entity.JadwalMakanan); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.JadwalMakanan)
		}
	}

	return arr
}

func (r *MockJadwalService) GetJadwalByIDService(id int) (entity.JadwalMakanan, error) {
	ret := r.Called(id)

	var jadwal entity.JadwalMakanan
	if res, ok := ret.Get(0).(func(int) entity.JadwalMakanan); ok {
		jadwal = res(id)
	} else {
		jadwal = ret.Get(0).(entity.JadwalMakanan)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return jadwal, err
}

func (r *MockJadwalService) UpdateJadwalByIDService(id int, jadwal entity.JadwalMakanan) error {
	ret := r.Called(id, jadwal)

	var err error
	if res, ok := ret.Get(0).(func(int, entity.JadwalMakanan) error); ok {
		err = res(id, jadwal)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockJadwalService) DeleteJadwalByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
