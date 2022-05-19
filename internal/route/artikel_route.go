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

func RegisterArtikelGroupAPI(e *echo.Echo, conf config.Config) {
	db := d.InitDB(conf)
	repo := r.NewMysqlArtikelRepository(db)
	service := u.NewServiceArtikel(repo, conf)
	cont := c.EchoArtikelController{Service: service}

	apiArtikel := e.Group("/artikel",
		middleware.Logger(),
		middleware.CORS(),
	)

	apiArtikel.GET("", cont.GetArtikelsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiArtikel.GET("/:id", cont.GetArtikelController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiArtikel.POST("", cont.CreateArtikelController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiArtikel.PUT("/:id", cont.UpdateArtikelController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiArtikel.DELETE("/:id", cont.DeleteArtikelController, middleware.JWT([]byte(conf.JWT_KEY)))
}
