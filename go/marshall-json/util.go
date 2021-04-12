package main

import (
	"encoding/json"
	"log"
)

func Json_parser(response string) ([]Vehicle, error) {
	var data []Vehicle

	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Fatal(err)
	}

	return data, err
}
