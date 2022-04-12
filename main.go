package main

import (
	"clean-arch/api"
	"clean-arch/app/modules"
	"clean-arch/config"
	"clean-arch/util"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.GetConfig()

	dbCon := util.NewConnectionDatabase(config)
	defer dbCon.CloseConnection()

	controllers := modules.RegisteModules(dbCon)

	e := echo.New()

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
