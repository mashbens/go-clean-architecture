package content

import (
	validator "github.com/go-playground/validator/v10"
)

type Repository interface {
	FindContentByID(id int) (content *Content, err error)
	FindAll() (contents []Content, err error)
	InserContent(content Content) (err error)
	PutContent(content Content) (err error)
}

type Service interface {
	GetContentByID(id int) (content *Content, err error)
	GetContents() (contents []Content, err error)
	CreateContent(content Content) (err error)
	UpdateContent(content Content) (err error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) GetContentByID(id int) (content *Content, err error) {
	result, err := s.repository.FindContentByID(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) GetContents() (contents []Content, err error) {
	result, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *service) CreateContent(content Content) (err error) {
	// validate masih error
	err = s.validate.Struct(content)
	if err != nil {
		return err
	}
	err = s.repository.InserContent(content)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateContent(content Content) (err error) {
	err = s.repository.PutContent(content)
	if err != nil {
		return err
	}
	return nil
}
