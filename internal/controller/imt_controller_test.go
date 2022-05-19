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

func TestCreateIMTController(t *testing.T) {
	service := imitation.MockIMTService{}

	imtController := EchoIMTController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error", func(t *testing.T) {
		service.On("CreateIMTService", mock.Anything).
			Return(errors.New("new")).Once()

		r := httptest.NewRequest("POST", "/imt", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.CreateIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("CreateIMTService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("POST", "/imt", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.CreateIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestGetIMTsController(t *testing.T) {
	service := imitation.MockIMTService{}

	imtController := EchoIMTController{
		Service: &service,
	}
	e := echo.New()

	t.Run("success", func(t *testing.T) {
		service.On("GetAllIMTService").
			Return([]entity.IMT{}).Once()

		r := httptest.NewRequest("GET", "/imt", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.GetIMTsController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestGetIMTController(t *testing.T) {
	service := imitation.MockIMTService{}

	imtController := EchoIMTController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("GetIMTByIDService", mock.Anything).
			Return(entity.IMT{}, errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("GET", "/imt/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.GetIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("GetIMTByIDService", mock.Anything).
			Return(entity.IMT{}, nil)

		r := httptest.NewRequest("GET", "/imt/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.GetIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestUpdateIMTController(t *testing.T) {
	service := imitation.MockIMTService{}

	imtController := EchoIMTController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("UpdateIMTByIDService", mock.Anything, mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("PUT", "/imt/:id", io.Reader(strings.NewReader(`{"Status": "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.UpdateIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("UpdateIMTByIDService", mock.Anything, mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("PUT", "/imt/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.UpdateIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
	})
}

func TestDeleteIMTController(t *testing.T) {
	service := imitation.MockIMTService{}

	imtController := EchoIMTController{
		Service: &service,
	}
	e := echo.New()

	t.Run("error not found", func(t *testing.T) {
		service.On("DeleteIMTByIDService", mock.Anything).
			Return(errors.New("Error Not Found")).Once()

		r := httptest.NewRequest("DELETE", "/imt/:id", io.Reader(strings.NewReader(`{"Status" : "Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.DeleteIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 404, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		service.On("DeleteIMTByIDService", mock.Anything).
			Return(nil).Once()

		r := httptest.NewRequest("DELETE", "/imt/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		err := imtController.DeleteIMTController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 204, w.Result().StatusCode)
	})
}
