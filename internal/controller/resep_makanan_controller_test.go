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

func TestCreateResepController(t *testing.T) {
	service := imitation.MockResepService{}

	resepController := EchoResepController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("CreateResepService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/resep-makanan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.CreateResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("CreateResepService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/resep-makanan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.CreateResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetResepsController(t *testing.T) {
	service := imitation.MockResepService{}

	resepController := EchoResepController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllResepService").
			Return([]entity.ResepMakanan{}).Once()

		r := httptest.NewRequest("GET", "/resep-makanan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.GetResepsController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetResepController(t *testing.T) {
	service := imitation.MockResepService{}

	resepController := EchoResepController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetResepByIDService", mock.Anything).
			Return(entity.ResepMakanan{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/resep-makanan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.GetResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetResepByIDService", mock.Anything).
			Return(entity.ResepMakanan{}, nil).Once()

		r := httptest.NewRequest("GET", "/resep-makanan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.GetResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateResepController(t *testing.T) {
	service := imitation.MockResepService{}

	resepController := EchoResepController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("UpdateResepByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/resep-makanan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.UpdateResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("UpdateResepByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/resep-makanan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.UpdateResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteResepController(t *testing.T) {
	service := imitation.MockResepService{}

	resepController := EchoResepController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteResepByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/resep-makanan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.DeleteResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("DeleteResepByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/resep-makanan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := resepController.DeleteResepController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
