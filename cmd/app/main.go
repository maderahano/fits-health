package main

import (
	"os"

	"github.com/labstack/echo/v4"

	conf "fits-health/config"
	docs "fits-health/docs"
	rest "fits-health/internal/route"

	echoSwag "github.com/swaggo/echo-swagger"
)

// @title Fits Health API Documentation
// @description This is Fits Health API
// @version 2.0
// @host localhost:8888
// @BasePath
// @schemes http https
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	config := conf.InitConf()
	e := echo.New()

	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterIMTGroupAPI(e, config)
	rest.RegisterProfilKesehatanGroupAPI(e, config)
	rest.RegisterResepMakananGroupAPI(e, config)
	rest.RegisterJadwalGroupAPI(e, config)
	rest.RegisterArtikelGroupAPI(e, config)
	rest.RegisterOlahragaGroupAPI(e, config)

	e.GET("/swagger/*", echoSwag.WrapHandler)
	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}
