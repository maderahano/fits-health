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

func TestCreateOlahragaController(t *testing.T) {
	service := imitation.MockOlahragaService{}

	olahragaController := EchoOlahragaController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("CreateOlahragaService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/olahraga", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.CreateOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("CreateOlahragaService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/olahraga", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.CreateOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetOlahragasController(t *testing.T) {
	service := imitation.MockOlahragaService{}

	olahragaController := EchoOlahragaController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllOlahragaService").
			Return([]entity.Olahraga{}).Once()

		r := httptest.NewRequest("GET", "/olahraga", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.GetOlahragasController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetOlahragaController(t *testing.T) {
	service := imitation.MockOlahragaService{}

	olahragaController := EchoOlahragaController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetOlahragaByIDService", mock.Anything).
			Return(entity.Olahraga{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/olahraga/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.GetOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetOlahragaByIDService", mock.Anything).
			Return(entity.Olahraga{}, nil).Once()

		r := httptest.NewRequest("GET", "/olahraga/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.GetOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateOlahragaController(t *testing.T) {
	service := imitation.MockOlahragaService{}

	olahragaController := EchoOlahragaController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("UpdateOlahragaByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/olahraga/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.UpdateOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("UpdateOlahragaByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/olahraga/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.UpdateOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteOlahragaController(t *testing.T) {
	service := imitation.MockOlahragaService{}

	olahragaController := EchoOlahragaController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteOlahragaByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/olahraga/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.DeleteOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("DeleteOlahragaByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/olahraga/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := olahragaController.DeleteOlahragaController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}
