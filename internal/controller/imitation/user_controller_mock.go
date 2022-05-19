package imitation

import (
	"fits-health/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (r *MockUserService) RegisterUserService(user entity.User) error {
	ret := r.Called(user)

	var err error
	if res, ok := ret.Get(0).(func(entity.User) error); ok {
		err = res(user)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockUserService) LoginUserService(email, password string) (string, int) {
	ret := r.Called(email, password)

	var strs string
	if res, ok := ret.Get(0).(func(string, string) string); ok {
		strs = res(email, password)
	} else {
		strs = ret.Get(0).(string)
	}

	var ints int
	if res, ok := ret.Get(1).(func(string, string) int); ok {
		ints = res(email, password)
	} else {
		ints = ret.Get(1).(int)
	}

	return strs, ints
}

func (r *MockUserService) GetAllUsersService() []entity.User {
	ret := r.Called()

	var arr []entity.User
	if res, ok := ret.Get(0).(func() []entity.User); ok {
		arr = res()
	} else {
		if ret.Get(0) != nil {
			arr = ret.Get(0).([]entity.User)
		}
	}

	return arr
}

func (r *MockUserService) GetUserByIDService(id int) (entity.User, error) {
	ret := r.Called(id)

	var user entity.User
	if res, ok := ret.Get(0).(func(int) entity.User); ok {
		user = res(id)
	} else {
		user = ret.Get(0).(entity.User)
	}

	var err error
	if res, ok := ret.Get(1).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(1)
	}

	return user, err
}

func (r *MockUserService) UpdateUserByIDService(id, idToken int, user entity.User) error {
	ret := r.Called(id, idToken, user)

	var err error
	if res, ok := ret.Get(0).(func(int, int, entity.User) error); ok {
		err = res(id, idToken, user)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (r *MockUserService) DeleteUserByIDService(id int) error {
	ret := r.Called(id)

	var err error
	if res, ok := ret.Get(0).(func(int) error); ok {
		err = res(id)
	} else {
		err = ret.Error(0)
	}

	return err
}

// type MockService struct {
// 	MFunc func(email, password string) (string, int)
// }

// func (m *MockService) SetFuncMock(f func(email, password string) (string, int)) {
// 	m.MFunc = f
// }

// func (m *MockService) RegisterUserService(user entity.User) error {
// 	return nil
// }

// func (m *MockService) LoginUserService(email, password string) (string, int) {
// 	return m.MFunc(email, password)
// }

// func (m *MockService) GetAllUsersService() []entity.User {
// 	return []entity.User{}
// }

// func (m *MockService) GetUserByIDService(id int) (entity.User, error) {
// 	return entity.User{}, nil
// }

// func (m *MockService) UpdateUserByIDService(id, idToken int, user entity.User) error {
// 	return nil
// }

// func (m *MockService) DeleteUserByIDService(id int) error {
// 	return nil
// }
