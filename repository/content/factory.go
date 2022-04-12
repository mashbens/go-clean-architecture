package content

import (
	"clean-arch/business/content"
	"clean-arch/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
	var contentRepository content.Repository

	if dbCon.Driver == util.MySQL {
		contentRepository = newMySQLRepository(dbCon.MySQL)
	} else {
		panic("Invalid database driver")
	}
	return contentRepository
}
