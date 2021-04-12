package product

import "github.com/mozarik/repository-pattern-simple/domain"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s Service) GetUser(id domain.ID) (*domain.Product, error) {
	return s.repo.Get(id)
}
