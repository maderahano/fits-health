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

func TestCreateArtikelController(t *testing.T) {
	service := imitation.MockArtikelService{}

	artikelController := EchoArtikelController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("CreateArtikelService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/artikel", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.CreateArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("CreateArtikelService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/artikel", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.CreateArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetArtikelsController(t *testing.T) {
	service := imitation.MockArtikelService{}

	artikelController := EchoArtikelController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllArtikelService").
			Return([]entity.Artikel{}).Once()

		r := httptest.NewRequest("GET", "/artikel", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.GetArtikelsController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetArtikelController(t *testing.T) {
	service := imitation.MockArtikelService{}

	artikelController := EchoArtikelController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetArtikelByIDService", mock.Anything).
			Return(entity.Artikel{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/artikel/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.GetArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetArtikelByIDService", mock.Anything).
			Return(entity.Artikel{}, nil)

		r := httptest.NewRequest("GET", "/artikel/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.GetArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateArtikelController(t *testing.T) {
	service := imitation.MockArtikelService{}

	artikelController := EchoArtikelController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("UpdateArtikelByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/artikel/:id", io.Reader(strings.NewReader(`{"Status": "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.UpdateArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("UpdateArtikelByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/artikel/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.UpdateArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteArtikelController(t *testing.T) {
	service := imitation.MockArtikelService{}

	artikelController := EchoArtikelController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteArtikelByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/artikel/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.DeleteArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("DeleteArtikelByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/artikel/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := artikelController.DeleteArtikelController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
