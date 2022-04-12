package modules

import (
	"clean-arch/api"
	contentV1Controller "clean-arch/api/v1/content"
	contentService "clean-arch/business/content"
	contentRepository "clean-arch/repository/content"
	"clean-arch/util"
)

func RegisteModules(dbCon *util.DatabaseConnection) api.Controller {
	contentPermitRepository := contentRepository.RepositoryFactory(dbCon)
	contentPermitService := contentService.NewService(contentPermitRepository)

	contentV1PermitController := contentV1Controller.NewController(contentPermitService)

	controllers := api.Controller{
		ContentV1Controller: contentV1PermitController,
	}
	return controllers
}
