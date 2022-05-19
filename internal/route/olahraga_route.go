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

func RegisterOlahragaGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlOlahragaRepository(db)
	service := u.NewServiceOlahraga(repo, conf)
	cont := c.EchoOlahragaController{Service: service}

	apiOlahraga := e.Group("/olahraga",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiOlahraga.GET("", cont.GetOlahragasController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiOlahraga.GET("/:id", cont.GetOlahragaController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiOlahraga.POST("", cont.CreateOlahragaController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiOlahraga.PUT("/:id", cont.UpdateOlahragaController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiOlahraga.DELETE("/:id", cont.DeleteOlahragaController, middleware.JWT([]byte(conf.JWT_KEY)))
}
