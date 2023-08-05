package product

import "time"

type Service interface {
	FindAll() ([]Product, error)
	Create(input CreateInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Product, error) {
	product, err := s.repository.FindAll()
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) Create(input CreateInput) (Product, error) {
	product := Product{}
	product.Name = input.Name
	product.Price = input.Price
	product.Quantity = input.Quantity
	product.Description = input.Description
	product.UserID = input.User.ID
	product.CategoryID = input.CategoryID

	now := time.Now()

	// Get the date by setting the time part to 00:00:00
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	product.Image = "https://img-global.cpcdn.com/recipes/93a46a53e22256b8/680x482cq70/songkolo-bagadang-ketan-serundeng-foto-resep-utama.jpg"
	product.EntryAt = today
	// product.ExpiredAt = input.ExpiredAt

	newProduct, err := s.repository.Create(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}
