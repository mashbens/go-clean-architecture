package content

import (
	"clean-arch/business/content"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func newMySQLRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repo *MySQLRepository) FindContentByID(id int) (content *content.Content, err error) {

	return content, nil
}

func (repo *MySQLRepository) FindAll() (contents []content.Content, err error) {
	result := repo.db.Find(&contents)
	if result.Error != nil {
		return nil, result.Error
	}
	return contents, nil
}

func (repo *MySQLRepository) InserContent(content content.Content) (err error) {

	return
}

func (repo *MySQLRepository) UpdateContent(content content.Content) (err error) {

	return
}
