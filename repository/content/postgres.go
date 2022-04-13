package content

import (
	"clean-arch/business/content"

	"gorm.io/gorm"
)

type PostgreSQLRepository struct {
	db *gorm.DB
}

func NewPostgreSQLRepository(db *gorm.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		db: db,
	}
}

func (repo *PostgreSQLRepository) FindContentByID(id int) (content *content.Content, err error) {

	return content, nil
}

func (repo *PostgreSQLRepository) FindAll() (contents []content.Content, err error) {
	result := repo.db.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil
}

func (repo *PostgreSQLRepository) InserContent(content content.Content) (err error) {

	return
}

func (repo *PostgreSQLRepository) UpdateContent(content content.Content) (err error) {

	return
}
