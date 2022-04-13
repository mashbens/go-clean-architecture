package content

import (
	"clean-arch/business/content"
	"clean-arch/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) content.Repository {
	var contentRepository content.Repository

	if dbCon.Driver == util.MySQL {
		contentRepository = NewMySQLRepository(dbCon.MySQL)
	} else if dbCon.Driver == util.PostgreSQL {
		contentRepository = NewPostgreSQLRepository(dbCon.PostgreSQL)
	}
	return contentRepository
}
