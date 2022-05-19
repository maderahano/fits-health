package controller

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoOlahragaController struct {
	Service domain.AdapterOlahragaService
}

// CreateOlahragaController godoc
// @Summary      Create Workout
// @Description  Developer can create workout for User
// @Tags         Olahraga
// @accept       json
// @Produce      json
// @Router       /olahraga [post]
// @param        data  body      entity.Olahraga  true  "required"
// @Success      200   {object}  entity.Olahraga
// @Failure      400   {object}  entity.Olahraga
// @Failure      500   {object}  entity.Olahraga
// @Security     JWT
func (ce *EchoOlahragaController) CreateOlahragaController(c echo.Context) error {
	olahraga := entity.Olahraga{}
	c.Bind(&olahraga)

	err := ce.Service.CreateOlahragaService(olahraga)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"olahraga": olahraga,
	})
}

// GetOlahragasController godoc
// @Summary      Get All Workout Information
// @Description  User can get information about all workout that can do it in home
// @Tags         Olahraga
// @accept       json
// @Produce      json
// @Router       /olahraga [get]
// @Success      200  {object}  entity.Olahraga
// @Failure      400  {object}  entity.Olahraga
// @Security     JWT
func (ce *EchoOlahragaController) GetOlahragasController(c echo.Context) error {
	olahragas := ce.Service.GetAllOlahragaService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"olahraga": olahragas,
	}, " ")
}

// GetOlahragaController godoc
// @Summary      Get Workout Information by ID
// @Description  User can get Workout Information by ID
// @Tags         Olahraga
// @accept       json
// @Produce      json
// @Router       /olahraga/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.Olahraga
// @Failure      400  {object}  entity.Olahraga
// @Security     JWT
func (ce *EchoOlahragaController) GetOlahragaController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetOlahragaByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"olahraga": res,
	})
}

// UpdateOlahragaController godoc
// @Summary      Update Olahraga
// @Description  Developer can update about workout information
// @Tags         Olahraga
// @accept       json
// @Produce      json
// @Router       /olahraga/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.Olahraga
// @Failure      400  {object}  entity.Olahraga
// @Failure      500  {object}  entity.Olahraga
// @Security     JWT
func (ce *EchoOlahragaController) UpdateOlahragaController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	olahraga := entity.Olahraga{}
	c.Bind(&olahraga)

	err := ce.Service.UpdateOlahragaByIDService(intID, olahraga)
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

// DeleteOlahragaController godoc
// @Summary      Delete Workout
// @Description  Developer can Delete Workout Information if workout information is useless
// @Tags         Olahraga
// @accept       json
// @Produce      json
// @Router       /olahraga/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.Olahraga
// @Failure      400  {object}  entity.Olahraga
// @Failure      500  {object}  entity.Olahraga
// @Security     JWT
func (ce *EchoOlahragaController) DeleteOlahragaController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteOlahragaByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "deleted",
		"id":       intID,
	})
}
