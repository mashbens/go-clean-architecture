package main

import (
	"clean-arch/api"
	"clean-arch/app/modules"
	"clean-arch/config"
	"clean-arch/util"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisteModules(dbCon, config)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		addres := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(addres); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
