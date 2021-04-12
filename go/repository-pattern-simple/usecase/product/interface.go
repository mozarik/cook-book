package product

import (
	"github.com/mozarik/repository-pattern-simple/domain"
)

type Reader interface {
	Get(id domain.ID) (*domain.Product, error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	GetUser(id domain.ID) (*domain.Product, error)
}
