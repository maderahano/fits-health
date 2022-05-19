package controller

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoIMTController struct {
	Service domain.AdapterIMTService
}

// CreateIMTController godoc
// @Summary      Create IMT for User
// @Description  IMT is for User that categorizing user body by weight and height
// @Tags         IMT
// @accept       json
// @Produce      json
// @Router       /imt [post]
// @param        data  body      entity.IMT  true  "required"
// @Success      200   {object}  entity.IMT
// @Failure      400   {object}  entity.IMT
// @Failure      500   {object}  entity.IMT
// @Security     JWT
func (ce *EchoIMTController) CreateIMTController(c echo.Context) error {
	imt := entity.IMT{}
	c.Bind(&imt)

	err := ce.Service.CreateIMTService(imt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"imt":      imt,
	})
}

// GetIMTsController godoc
// @Summary      Get All IMT Information
// @Description  From IMT information, user can measure itself by knowing what type of body he/she get from weight and height
// @Tags         IMT
// @accept       json
// @Produce      json
// @Router       /imt [get]
// @Success      200  {object}  entity.IMT
// @Failure      400  {object}  entity.IMT
// @Security     JWT
func (ce *EchoIMTController) GetIMTsController(c echo.Context) error {
	imts := ce.Service.GetAllIMTService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"imt":      imts,
	}, " ")
}

// GetIMTController godoc
// @Summary      Get IMT Information by ID
// @Description  User can get IMT Information by ID
// @Tags         IMT
// @accept       json
// @Produce      json
// @Router       /imt/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.IMT
// @Failure      400  {object}  entity.IMT
// @Security     JWT
func (ce *EchoIMTController) GetIMTController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetIMTByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"imt":      res,
	})
}

// UpdateIMTController godoc
// @Summary      Update IMT
// @Description  Developer can update about IMT information
// @Tags         IMT
// @accept       json
// @Produce      json
// @Router       /imt/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.IMT
// @Failure      400  {object}  entity.IMT
// @Failure      500  {object}  entity.IMT
// @Security     JWT
func (ce *EchoIMTController) UpdateIMTController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	imt := entity.IMT{}
	c.Bind(&imt)

	err := ce.Service.UpdateIMTByIDService(intID, imt)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id":       intID,
	})
}

// DeleteIMTController godoc
// @Summary      Delete IMT
// @Description  Developer can Delete IMT Information if IMT information is wrong
// @Tags         IMT
// @accept       json
// @Produce      json
// @Router       /imt/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.IMT
// @Failure      400  {object}  entity.IMT
// @Failure      500  {object}  entity.IMT
// @Security     JWT
func (ce *EchoIMTController) DeleteIMTController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteIMTByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
