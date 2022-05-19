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

func RegisterResepMakananGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlResepRepository(db)
	service := u.NewServiceResep(repo, conf)
	cont := c.EchoResepController{Service: service}

	apiResep := e.Group("resep-makanan",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiResep.GET("", cont.GetResepsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiResep.GET("/:id", cont.GetResepController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiResep.POST("", cont.CreateResepController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiResep.PUT("/:id", cont.UpdateResepController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiResep.DELETE(":/id", cont.DeleteResepController, middleware.JWT([]byte(conf.JWT_KEY)))
}
