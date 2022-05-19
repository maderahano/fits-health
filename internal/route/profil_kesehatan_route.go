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

func RegisterProfilKesehatanGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlProfilKesehatanRepository(db)
	service := u.NewServiceProfileKesehatan(repo, conf)
	cont := c.EchoProfilKesehatanController{Service: service}

	apiProfilKesehatan := e.Group("/profil-kesehatan",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiProfilKesehatan.GET("", cont.GetProfilKesehatansController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiProfilKesehatan.GET("/:id", cont.GetProfilKesehatanController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiProfilKesehatan.POST("", cont.CreateProfilKesehatanController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiProfilKesehatan.PUT("/:id", cont.UpdateProfilKesehatanByIDController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiProfilKesehatan.DELETE("/:id", cont.DeleteProfilKesehatanByIDController, middleware.JWT([]byte(conf.JWT_KEY)))
}
