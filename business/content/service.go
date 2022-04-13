package content

type Repository interface {
	FindContentByID(id int) (content *Content, err error)
	FindAll() (contents []Content, err error)
	InserContent(content Content) (err error)
	UpdateContent(content Content) (err error)
}

type Service interface {
	GetContentByID(id int) (content *Content, err error)
	GetContents() (contents []Content, err error)
	CreateContent(content Content) (err error)
	UpdateContent(content Content) (err error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
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
	return
}

func (s *service) UpdateContent(content Content) (err error) {
	return
}
