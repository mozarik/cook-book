package main

import (
	"fmt"

	"github.com/mozarik/repository-pattern-simple/repository"
)

func main() {
	user := repository.NewUserInMem()

	data, _ := user.Get()

	fmt.Print(data)
}
