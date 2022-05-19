package controller

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoProfilKesehatanController struct {
	Service domain.AdapterProfilKesehatanService
}

// CreateProfilKesehatanController godoc
// @Summary      Create Health Information for user
// @Description  Health Information is about user personal health information
// @Tags         ProfilKesehatan
// @accept       json
// @Produce      json
// @Router       /profil-kesehatan [post]
// @param        data  body      entity.ProfilKesehatan  true  "required"
// @Success      200   {object}  entity.ProfilKesehatan
// @Failure      400   {object}  entity.ProfilKesehatan
// @Failure      500   {object}  entity.ProfilKesehatan
// @Security     JWT
func (ce *EchoProfilKesehatanController) CreateProfilKesehatanController(c echo.Context) error {
	profil := entity.ProfilKesehatan{}
	c.Bind(&profil)

	err := ce.Service.CreateProfilKesehatanService(profil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages":         "success",
		"profil kesehatan": profil,
	})
}

// GetProfilKesehatansController godoc
// @Summary      Get All Health Information for developer
// @Description  Developer can get all health information from user
// @Tags         ProfilKesehatan
// @accept       json
// @Produce      json
// @Router       /profil-kesehatan [get]
// @Success      200  {object}  entity.ProfilKesehatan
// @Failure      400  {object}  entity.ProfilKesehatan
// @Security     JWT
func (ce *EchoProfilKesehatanController) GetProfilKesehatansController(c echo.Context) error {
	profils := ce.Service.GetAllProfilKesehatanService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages":         "success",
		"profil kesehatan": profils,
	}, " ")
}

// GetIMTController godoc
// @Summary      Get Health Information by ID
// @Description  Developer and User can get Health Information by ID
// @Tags         ProfilKesehatan
// @accept       json
// @Produce      json
// @Router       /profil-kesehatan/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.ProfilKesehatan
// @Failure      400  {object}  entity.ProfilKesehatan
// @Security     JWT
func (ce *EchoProfilKesehatanController) GetProfilKesehatanController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetProfilKesehatanByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":         "success",
		"profil kesehatan": res,
	})
}

// UpdateProfilKesehatanByIDController godoc
// @Summary      Update Health Information
// @Description  User can update about Health information
// @Tags         ProfilKesehatan
// @accept       json
// @Produce      json
// @Router       /profil-kesehatan/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.ProfilKesehatan
// @Failure      400  {object}  entity.ProfilKesehatan
// @Failure      500  {object}  entity.ProfilKesehatan
// @Security     JWT
func (ce *EchoProfilKesehatanController) UpdateProfilKesehatanByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	profil := entity.ProfilKesehatan{}
	c.Bind(&profil)

	err := ce.Service.UpdateProfilKesehatanByIDService(intID, profil)
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

// DeleteProfilKesehatanByIDController godoc
// @Summary      Delete Health Information
// @Description  Developer can Delete Health Information if Health information is wrong
// @Tags         ProfilKesehatan
// @accept       json
// @Produce      json
// @Router       /profil-kesehatan/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.ProfilKesehatan
// @Failure      400  {object}  entity.ProfilKesehatan
// @Failure      500  {object}  entity.ProfilKesehatan
// @Security     JWT
func (ce *EchoProfilKesehatanController) DeleteProfilKesehatanByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteProfilKesehatanByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
