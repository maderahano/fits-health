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

func TestCreateJadwalController(t *testing.T) {
	service := imitation.MockJadwalService{}

	jadwalController := EchoJadwalController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("CreateJadwalService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/jadwal-makanan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.CreateJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("CreateJadwalService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/jadwal-makanan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.CreateJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetJadwalsController(t *testing.T) {
	service := imitation.MockJadwalService{}

	jadwalController := EchoJadwalController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllJadwalService").
			Return([]entity.JadwalMakanan{}).Once()

		r := httptest.NewRequest("GET", "/jadwal-makanan", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.GetJadwalsController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetJadwalController(t *testing.T) {
	service := imitation.MockJadwalService{}

	jadwalController := EchoJadwalController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetJadwalByIDService", mock.Anything).
			Return(entity.JadwalMakanan{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/jadwal-makanan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.GetJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetJadwalByIDService", mock.Anything).
			Return(entity.JadwalMakanan{}, nil).Once()

		r := httptest.NewRequest("GET", "/jadwal-makanan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.GetJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateJadwalController(t *testing.T) {
	service := imitation.MockJadwalService{}

	jadwalController := EchoJadwalController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("UpdateJadwalByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/jadwal-makanan/:id", io.Reader(strings.NewReader(`{"Status": "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.UpdateJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("UpdateJadwalByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/jadwal-makanan/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.UpdateJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteJadwalController(t *testing.T) {
	service := imitation.MockJadwalService{}

	jadwalController := EchoJadwalController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteJadwalByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/jadwal-makanan/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := jadwalController.DeleteJadwalController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})
}
