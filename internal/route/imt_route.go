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

func RegisterIMTGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlIMTRepository(db)
	service := u.NewServiceIMT(repo, conf)
	cont := c.EchoIMTController{Service: service}

	apiIMT := e.Group("/imt",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiIMT.GET("", cont.GetIMTsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiIMT.GET("/:id", cont.GetIMTController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiIMT.POST("", cont.CreateIMTController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiIMT.PUT("/:id", cont.UpdateIMTController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiIMT.DELETE("/:id", cont.DeleteIMTController, middleware.JWT([]byte(conf.JWT_KEY)))
}
