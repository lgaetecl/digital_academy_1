package products

type Service interface {
	Create(product Product) error
	FindById(id int32) (*Product, error)
	FindAll() *[]Product
	Delete(id int32) error
	Update(id int32, product Product) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(product Product) error {
	return nil
}

func (s *service) FindById(id int32) (*Product, error) {
	return nil, nil
}

func (s *service) FindAll() *[]Product {
	return nil
}

func (s *service) Delete(id int32) error {
	return s.repo.Delete(id)
}

func (s *service) Update(id int32, products Product) (*Product, error) {
	return nil, nil
}
