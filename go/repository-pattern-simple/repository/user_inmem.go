package repository

import (
	"encoding/json"
	"log"
	"os"

	"github.com/mozarik/repository-pattern-simple/domain"
)

type UserInMem struct{}

func NewUserInMem() *UserInMem {
	return &UserInMem{}
}

func (r *UserInMem) Get() (domain.Product, error) {
	return getData()
}

// getData(id) return domain.Product
func getData() (domain.Product, error) {
	data, err := Json_parser()
	if err != nil {
		log.Fatal(err)
	}

	return data, err
}

func Json_parser() (domain.Product, error) {
	file, _ := os.ReadFile("p1.json")
	var data domain.Product

	err := json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data, err
}
