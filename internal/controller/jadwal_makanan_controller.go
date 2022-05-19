package controller

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoJadwalController struct {
	Service domain.AdapterJadwalService
}

// CreateJadwalController godoc
// @Summary      Create Food Schedule
// @Description  User can create Food Schedule everyday in their life
// @Tags         JadwalMakanan
// @accept       json
// @Produce      json
// @Router       /jadwal-makanan [post]
// @param        data  body      entity.JadwalMakanan  true  "required"
// @Success      200   {object}  entity.JadwalMakanan
// @Failure      400   {object}  entity.JadwalMakanan
// @Failure      500   {object}  entity.JadwalMakanan
// @Security     JWT
func (ce *EchoJadwalController) CreateJadwalController(c echo.Context) error {
	jadwal := entity.JadwalMakanan{}
	c.Bind(&jadwal)

	err := ce.Service.CreateJadwalService(jadwal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"imt":      jadwal,
	})
}

// GetJadwalsController godoc
// @Summary      Get All Food Schedule Information in User
// @Description  User can get their food schedule information
// @Tags         JadwalMakanan
// @accept       json
// @Produce      json
// @Router       /jadwal-makanan [get]
// @Success      200  {object}  entity.JadwalMakanan
// @Failure      400  {object}  entity.JadwalMakanan
// @Security     JWT
func (ce *EchoJadwalController) GetJadwalsController(c echo.Context) error {
	jadwals := ce.Service.GetAllJadwalService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"jadwal":   jadwals,
	}, " ")
}

// GetJadwalController godoc
// @Summary      Get Food Schedule Information by ID
// @Description  User can get Food Schedule Information by ID
// @Tags         JadwalMakanan
// @accept       json
// @Produce      json
// @Router       /jadwal-makanan/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.JadwalMakanan
// @Failure      400  {object}  entity.JadwalMakanan
// @Security     JWT
func (ce *EchoJadwalController) GetJadwalController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetJadwalByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"jadwal":   res,
	})
}

// UpdateJadwalController godoc
// @Summary      Update Food Schedule
// @Description  User can update about Food Schedule if their schdule got wrong
// @Tags         JadwalMakanan
// @accept       json
// @Produce      json
// @Router       /jadwal-makanan/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.JadwalMakanan
// @Failure      400  {object}  entity.JadwalMakanan
// @Failure      500  {object}  entity.JadwalMakanan
// @Security     JWT
func (ce *EchoJadwalController) UpdateJadwalController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	jadwal := entity.JadwalMakanan{}
	c.Bind(&jadwal)

	err := ce.Service.UpdateJadwalByIDService(intID, jadwal)
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

// DeleteJadwalController godoc
// @Summary      Delete Food Schedule
// @Description  Developer can Delete Food Schedule Information if they don't like about their food schedule that they created
// @Tags         JadwalMakanan
// @accept       json
// @Produce      json
// @Router       /jadwal-makanan/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.JadwalMakanan
// @Failure      400  {object}  entity.JadwalMakanan
// @Failure      500  {object}  entity.JadwalMakanan
// @Security     JWT
func (ce *EchoJadwalController) DeleteJadwalController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteJadwalByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
