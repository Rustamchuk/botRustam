package product

import "fmt"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(ind int) (*Product, error) {
	if ind < 0 || ind >= len(allProducts) {
		return nil, fmt.Errorf("wrong IND")
	}
	return &allProducts[ind], nil
}
