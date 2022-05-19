package controller

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoResepController struct {
	Service domain.AdapterResepService
}

// CreateResepController godoc
// @Summary      Create Food Recipe for User
// @Description  Food Recipe is for user that user can use the food recipe for cooking or for food schedule
// @Tags         ResepMakanan
// @accept       json
// @Produce      json
// @Router       /resep-makanan [post]
// @param        data  body      entity.ResepMakanan  true  "required"
// @Success      200   {object}  entity.ResepMakanan
// @Failure      400   {object}  entity.ResepMakanan
// @Failure      500   {object}  entity.ResepMakanan
// @Security     JWT
func (ce *EchoResepController) CreateResepController(c echo.Context) error {
	resep := entity.ResepMakanan{}
	c.Bind(&resep)

	err := ce.Service.CreateResepService(resep)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"imt":      resep,
	})
}

// GetResepsController godoc
// @Summary      Get All Food Recipe Information
// @Description  User can get information about food recipe that can be useful in their life
// @Tags         ResepMakanan
// @accept       json
// @Produce      json
// @Router       /resep-makanan [get]
// @Success      200  {object}  entity.ResepMakanan
// @Failure      400  {object}  entity.ResepMakanan
// @Security     JWT
func (ce *EchoResepController) GetResepsController(c echo.Context) error {
	reseps := ce.Service.GetAllResepService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"resep":    reseps,
	}, " ")
}

// GetResepController godoc
// @Summary      Get IMT Food Recipe by ID
// @Description  User can get Food Recipe Information by ID
// @Tags         ResepMakanan
// @accept       json
// @Produce      json
// @Router       /resep-makanan/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.ResepMakanan
// @Failure      400  {object}  entity.ResepMakanan
// @Security     JWT
func (ce *EchoResepController) GetResepController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetResepByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"resep":    res,
	})
}

// UpdateResepController godoc
// @Summary      Update Food Recipe
// @Description  Developer can update about food recipe information if the information about food recipe is wrong
// @Tags         ResepMakanan
// @accept       json
// @Produce      json
// @Router       /resep-makanan/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.ResepMakanan
// @Failure      400  {object}  entity.ResepMakanan
// @Failure      500  {object}  entity.ResepMakanan
// @Security     JWT
func (ce *EchoResepController) UpdateResepController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resep := entity.ResepMakanan{}
	c.Bind(&resep)

	err := ce.Service.UpdateResepByIDService(intID, resep)
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

// DeleteResepController godoc
// @Summary      Delete Food Recipe
// @Description  Developer can Delete Food Recipe Information if Food Recipe information is useless
// @Tags         ResepMakanan
// @accept       json
// @Produce      json
// @Router       /resep-makanan/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.ResepMakanan
// @Failure      400  {object}  entity.ResepMakanan
// @Failure      500  {object}  entity.ResepMakanan
// @Security     JWT
func (ce *EchoResepController) DeleteResepController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteResepByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
