package modules

import (
	"clean-arch/api"
	contentV1Controller "clean-arch/api/v1/content"
	contentService "clean-arch/business/content"
	contentRepository "clean-arch/repository/content"

	"clean-arch/config"

	authController "clean-arch/api/v1/auth"
	authService "clean-arch/business/auth"

	"clean-arch/util"
)

func RegisteModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	contentPermitRepository := contentRepository.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	authPermitService := authService.NewService(config)
	authPermitController := authController.NewController(authPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
		AuthController:      authPermitController,
	}
	return controllers
}
