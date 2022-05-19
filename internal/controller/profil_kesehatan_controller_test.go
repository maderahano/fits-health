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

func TestCreateProfilKesehatanController(t *testing.T) {
	service := imitation.MockProfilService{}

	profilController := EchoProfilKesehatanController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("CreateProfilKesehatanService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/profil-kesehatan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.CreateProfilKesehatanController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("CreateProfilKesehatanService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/profil-kesehatan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.CreateProfilKesehatanController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetProfilKesehatansController(t *testing.T) {
	service := imitation.MockProfilService{}

	profilController := EchoProfilKesehatanController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllProfilKesehatanService").
			Return([]entity.ProfilKesehatan{}).Once()

		r := httptest.NewRequest("GET", "/profil-kesehatan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.GetProfilKesehatansController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetProfilKesehatanController(t *testing.T) {
	service := imitation.MockProfilService{}

	profilController := EchoProfilKesehatanController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetProfilKesehatanByIDService", mock.Anything).
			Return(entity.ProfilKesehatan{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/profil-kesehatan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.GetProfilKesehatanController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetProfilKesehatanByIDService", mock.Anything).
			Return(entity.ProfilKesehatan{}, nil).Once()

		r := httptest.NewRequest("GET", "/profil-kesehatan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.GetProfilKesehatanController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateProfilKesehatanByIDController(t *testing.T) {
	service := imitation.MockProfilService{}

	profilController := EchoProfilKesehatanController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("UpdateProfilKesehatanByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/profil-kesehatan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.UpdateProfilKesehatanByIDController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("UpdateProfilKesehatanByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/profil-kesehatan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.UpdateProfilKesehatanByIDController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteProfilKesehatanByIDController(t *testing.T) {
	service := imitation.MockProfilService{}

	profilController := EchoProfilKesehatanController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteProfilKesehatanByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/profil-kesehatan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.DeleteProfilKesehatanByIDController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("DeleteProfilKesehatanByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/profil-kesehatan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := profilController.DeleteProfilKesehatanByIDController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
