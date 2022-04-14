package content

import (
	"clean-arch/business/content"

	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (repo *MySQLRepository) FindContentByID(id int) (content *content.Content, err error) {
	result := repo.db.First(&content, id)
	if result.Error != nil {
		return nil, result.Error
	}
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
	result := repo.db.Create(&content)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *MySQLRepository) PutContent(content content.Content) (err error) {
	result := repo.db.Save(&content)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
