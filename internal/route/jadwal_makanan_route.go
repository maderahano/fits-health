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

func RegisterJadwalGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlJadwalRepository(db)
	service := u.NewServiceJadwal(repo, conf)
	cont := c.EchoJadwalController{Service: service}

	apiJadwal := e.Group("/jadwal-makanan",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiJadwal.GET("", cont.GetJadwalsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiJadwal.GET("/:id", cont.GetJadwalController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiJadwal.POST("", cont.CreateJadwalController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiJadwal.PUT("/:id", cont.UpdateJadwalController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiJadwal.DELETE("/:id", cont.DeleteJadwalController, middleware.JWT([]byte(conf.JWT_KEY)))
}
