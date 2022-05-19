package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"fits-health/internal/domain"
	"fits-health/internal/entity"
)

type EchoUserController struct {
	Service domain.AdapterUserService
}

// RegisterUserController godoc
// @Summary      Register User
// @Description  People can Register as a User
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /register [post]
// @param        data  body      entity.User  true  "required"
// @Success      200   {object}  entity.User
// @Failure      400   {object}  entity.User
// @Failure      500   {object}  entity.User
func (ce *EchoUserController) RegisterUserController(c echo.Context) error {
	user := entity.User{}
	c.Bind(&user)

	err := ce.Service.RegisterUserService(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"users":    user,
	})
}

// LoginUserController godoc
// @Summary      Login User
// @Description  Login User
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /login [post]
// @Param        data  body      entity.User  true  "required"
// @Success      200   {object}  entity.User
// @Success      400   {object}  entity.User
// @Failure      401   {object}  entity.User
// @Failure      500   {object}  entity.User
func (ce *EchoUserController) LoginUserController(c echo.Context) error {
	userLogin := make(map[string]interface{})
	c.Bind(&userLogin)

	token, statusCode := ce.Service.LoginUserService(userLogin["email"].(string), userLogin["password"].(string))
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "email atau password salah",
		}, " ")
	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal server error",
		}, " ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, " ")
}

// GetUserController godoc
// @Summary      Get User By Id
// @Description  User Can Get User Information By Id
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /users/{id} [get]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.User
// @Failure      400  {object}  entity.User
// @Security     JWT
func (ce *EchoUserController) GetUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	res, err := ce.Service.GetUserByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    res,
	})
}

// GetUsersController godoc
// @Summary      Get All User
// @Description  Get ALl User
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /users [get]
// @Success      200  {object}  entity.User
// @Failure      400  {object}  entity.User
// @Security     JWT
func (ce *EchoUserController) GetUsersController(c echo.Context) error {
	users := ce.Service.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	}, " ")
}

// UpdateUserController godoc
// @Summary      Update User
// @Description  User can Update their status or information
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /users/{id} [put]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.User
// @Failure      400  {object}  entity.User
// @Failure      500  {object}  entity.User
// @Security     JWT
func (ce *EchoUserController) UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := entity.User{}
	c.Bind(&user)

	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	err := ce.Service.UpdateUserByIDService(intID, int(claim["id"].(float64)), user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":    "edited",
		"id":          intID,
		"expected id": claim["id"],
	})
}

// DeleteUserController godoc
// @Summary      Delete User
// @Description  User can Delete their own account
// @Tags         User
// @accept       json
// @Produce      json
// @Router       /users/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entity.User
// @Failure      400  {object}  entity.User
// @Failure      500  {object}  entity.User
// @Security     JWT
func (ce *EchoUserController) DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.Service.DeleteUserByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}
