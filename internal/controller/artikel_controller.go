package controller

import (
	"fits-health/internal/domain"
	"fits-health/internal/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EchoArtikelController struct {
	Service domain.AdapterArtikelService
}

// CreateArtikelController godoc
// @Summary      Create Article
// @Description  Developer can Create Article about healthy, disease, and many more
// @Tags         Artikel
// @accept       json
// @Produce      json
// @Router       /artikel [post]
// @param        data  body      entity.Artikel  true  "required"
// @Success      200   {object}  entity.Artikel
// @Failure      400   {object}  entity.Artikel
// @Failure      500   {object}  entity.Artikel
// @Security     JWT
func (ce *EchoArtikelController) CreateArtikelController(c echo.Context) error {
	artikel := entity.Artikel{}
	c.Bind(&artikel)

	err := ce.Service.CreateArtikelService(artikel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"imt":      artikel,
	})
}

// GetArtikelsController godoc
// @Summary      Get All Artikel Information
// @Description  User can Get Information article about health that posting by admin or developer
// @Tags         Artikel
// @accept       json
// @Produce      json
// @Router       /artikel [get]
// @Success      200  {object}  entity.Artikel
// @Failure      400  {object}  entity.Artikel
// @Security     JWT
func (ce *EchoArtikelController) GetArtikelsController(c echo.Context) error {
	artikels := ce.Service.GetAllArtikelService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"artikel":  artikels,
	}, " ")
}

// GetArtikelController godoc
// @Summary      Get Article Information by ID
// @Description  User can get Article Information by ID
// @Tags         Artikel
// @accept       json
// @Produce      json
// @Router       /artikel/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.Artikel
// @Failure      400  {object}  entity.Artikel
// @Security     JWT
func (ce *EchoArtikelController) GetArtikelController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetArtikelByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"artikel":  res,
	})
}

// UpdateArtikelController godoc
// @Summary      Update Artikel
// @Description  Developer can update about article information that if the information article is wrong
// @Tags         Artikel
// @accept       json
// @Produce      json
// @Router       /artikel/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.Artikel
// @Failure      400  {object}  entity.Artikel
// @Failure      500  {object}  entity.Artikel
// @Security     JWT
func (ce *EchoArtikelController) UpdateArtikelController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	artikel := entity.Artikel{}
	c.Bind(&artikel)

	err := ce.Service.UpdateArtikelByIDService(intID, artikel)
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

// DeleteArtikelController godoc
// @Summary      Delete Artikel
// @Description  Developer can Delete Article Information if they want to delete it
// @Tags         Artikel
// @accept       json
// @Produce      json
// @Router       /artikel/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.Artikel
// @Failure      400  {object}  entity.Artikel
// @Failure      500  {object}  entity.Artikel
// @Security     JWT
func (ce *EchoArtikelController) DeleteArtikelController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteArtikelByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
