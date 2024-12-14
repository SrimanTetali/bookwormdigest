package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

//a function that reads a file and returns a list of bookworms

func loadBookworms(filepath string) ([]Bookworm, error) {

	// open file using os package
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// decoding json file
	var bookworms []Bookworm
	err = json.NewDecoder(file).Decode(&bookworms)
	if err != nil {
		return nil, err
	}

	return bookworms, nil

}
