package controller

import (
	"errors"
	"fits-health/internal/controller/imitation"
	"fits-health/internal/entity"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUserController(t *testing.T) {
	service := imitation.MockUserService{}

	userController := EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("RegisterUserService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.RegisterUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("RegisterUserService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/register", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.RegisterUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

// func TestLoginUserController(t *testing.T) {
// 	service := imitation.MockUserService{}
// 	userController := EchoUserController{
// 		Service: &service,
// 	}
// 	e := echo.New()

// 	t.Run("error", func(t *testing.T) {
// 		service.On("LoginUserService", mock.Anything, mock.Anything).
// 			Return("Error", 500).Once()

// 		r := httptest.NewRequest("POST", "/login", io.Reader(strings.NewReader(`{"Status": "Error"}`)))
// 		w := httptest.NewRecorder()
// 		echoContext := e.NewContext(r, w)

// 		err := userController.LoginUserController(echoContext)
// 		if err != nil {
// 			return
// 		}

// 		assert.Equal(t, 500, w.Result().StatusCode)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		service.On("LoginUserService", mock.Anything, mock.Anything).
// 			Return("Success", 200).Once()

// 		bodyReader := strings.NewReader("{'Email': '12124', 'Password': 'testinasg'}")
// 		r := httptest.NewRequest("POST", "/login", bodyReader)
// 		w := httptest.NewRecorder()
// 		echoContext := e.NewContext(r, w)

// 		err := userController.LoginUserController(echoContext)
// 		if err != nil {
// 			return
// 		}

// 		assert.Equal(t, 200, w.Result().StatusCode)
// 	})
// }

// func TestLoginUserController(t *testing.T) {
// 	service := imitation.MockUserService{}
// 	mockLogin := new(imitation.MockUserService)
// 	loginData := entity.User{
// 		Email:    "mdrahano@gmail.com",
// 		Password: "luwakwaitkopisemakindidepan",
// 	}

// 	userController := EchoUserController{
// 		Service: &service,
// 	}
// 	e := echo.New()

// 	t.Run("error", func(t *testing.T) {
// 		mockLogin.On("LoginUserService", mock.Anything, mock.Anything).
// 			Return("Error", 500).Once()
// 		mockLogin.LoginUserService(loginData.Email, loginData.Password)
// 		service.On("LoginUserService", mock.Anything, mock.Anything).
// 			Return("Error", 500).Once()

// 		r := httptest.NewRequest("POST", "/login", io.Reader(strings.NewReader(`{"Status": "Error"}`)))
// 		w := httptest.NewRecorder()
// 		echoContext := e.NewContext(r, w)

// 		err := userController.LoginUserController(echoContext)
// 		if err != nil {
// 			return
// 		}

// 		assert.Equal(t, 500, w.Result().StatusCode)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		mockLogin.On("LoginUserService", mock.Anything, mock.Anything).
// 			Return("Success", 200).Once()
// 		mockLogin.LoginUserService(loginData.Email, loginData.Password)
// 		service.On("LoginUserService", mock.Anything, mock.Anything).
// 			Return("Success", 200).Once()

// 		r := httptest.NewRequest("POST", "/login", nil)
// 		w := httptest.NewRecorder()
// 		echoContext := e.NewContext(r, w)

// 		err := userController.LoginUserController(echoContext)
// 		if err != nil {
// 			return
// 		}

// 		assert.Equal(t, 200, w.Result().StatusCode)
// 	})
// }

func TestGetAllUsersService(t *testing.T) {
	service := imitation.MockUserService{}

	userController := EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllUsersService").
			Return([]entity.User{}).Once()

		r := httptest.NewRequest("GET", "/users", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUsersController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetUserByIDService(t *testing.T) {
	service := imitation.MockUserService{}

	userController := EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetUserByIDService", mock.Anything).
			Return(entity.User{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/users/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetUserByIDService", mock.Anything).
			Return(entity.User{}, nil).Once()
		r := httptest.NewRequest("GET", "/users/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.GetUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

// func TestUpdateUserController(t *testing.T) {
// 	service := imitation.MockUserService{}

// 	userController := EchoUserController{
// 		Service: &service,
// 	}
// 	e := echo.New()

// 	t.Run("error", func(t *testing.T) {
// 		service.On("UpdateUserByIDService", mock.Anything, mock.Anything, mock.Anything).
// 			Return(errors.New("Error")).Once()

// 		r := httptest.NewRequest("PUT", "/users/:id", io.Reader(strings.NewReader(`{"Status": "Error"}`)))
// 		w := httptest.NewRecorder()
// 		echoContext := e.NewContext(r, w)

// 		err := userController.UpdateUserController(echoContext)
// 		if err != nil {
// 			return
// 		}

// 		assert.Equal(t, 500, w.Result().StatusCode)
// 	})

// 	t.Run("success", func(t *testing.T) {
// 		service.On("UpdateUserByIDService", mock.Anything, mock.Anything, mock.Anything).
// 			Return(nil).Once()

// 		r := httptest.NewRequest("PUT", "/users/:id", nil)
// 		w := httptest.NewRecorder()
// 		echoContext := e.NewContext(r, w)

// 		err := userController.UpdateUserController(echoContext)
// 		if err != nil {
// 			return
// 		}

// 		assert.Equal(t, 200, w.Result().StatusCode)
// 	})
// }

func TestDeleteUserController(t *testing.T) {
	service := imitation.MockUserService{}

	userController := EchoUserController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteUserByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/users/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.DeleteUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("DeleteUserByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/users/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := userController.DeleteUserController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
