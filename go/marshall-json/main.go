package main

import (
	"fmt"
	"log"
)

type Vehicle struct {
	Type   string
	Wheels int
}

func main() {
	response := `[{"type": "Car","wheels": 4},{"type": "Motorcycle","wheels": 2}]`

	data, err := Json_parser(response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(data)
	fmt.Print("Success")
}
