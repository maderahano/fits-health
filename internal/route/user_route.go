package route

import (
	"fits-health/config"
	d "fits-health/database"
	c "fits-health/internal/controller"
	r "fits-health/internal/repository"
	u "fits-health/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlUserRepository(db)
	service := u.NewServiceUser(repo, conf)
	cont := c.EchoUserController{Service: service}

	e.POST("/login", cont.LoginUserController, middleware.Logger(), middleware.CORS())
	e.POST("/register", cont.RegisterUserController, middleware.Logger(), middleware.CORS())

	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
		// m.APIKEYMiddleware,
	)

	apiUser.GET("", cont.GetUsersController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token tidak valid",
				}, " ")
			},
			SuccessHandler: func(c echo.Context) {},
		},
	))
	apiUser.GET("/:id", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/:id", cont.UpdateUserController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiUser.DELETE("/:id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))
}
