package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockProfilService struct {
	mock.Mock
}

func (r *MockProfilService) CreateProfilKesehatanService(profil entity.ProfilKesehatan) error {
	ret := r.Called(profil)

	var err error
	if res, ok := ret.Get(0).(func(entity.ProfilKesehatan) error); ok {
		err = res(profil)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockProfilService) GetAllProfilKesehatanService() []entity.ProfilKesehatan {
	ret := r.Called()

	var arr []entity.ProfilKesehatan
	if res, ok := ret.Get(0).(func() []entity.ProfilKesehatan); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.ProfilKesehatan)
		}
	}

	return arr
}

func (r *MockProfilService) GetProfilKesehatanByIDService(id int) (entity.ProfilKesehatan, error) {
	ret := r.Called(id)

	var profil entity.ProfilKesehatan
	if res, ok := ret.Get(0).(func(int) entity.ProfilKesehatan); ok {
		profil = res(id)
	} else {
		profil = ret.Get(0).(entity.ProfilKesehatan)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return profil, err
}

func (r *MockProfilService) UpdateProfilKesehatanByIDService(id int, profil entity.ProfilKesehatan) error {
	ret := r.Called(id, profil)

	var err error
	if res, ok := ret.Get(0).(func(int, entity.ProfilKesehatan) error); ok {
		err = res(id, profil)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockProfilService) DeleteProfilKesehatanByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}
