package main

import (
    "encoding/json"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// loadBookworms reads the file and returns the list of bookworms
// and their beloved books, found therein.
func loadBookworms(filepath string) ([]Bookworm, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

    var bookworms []Bookworm
    err = json.NewDecoder(f).Decode(&bookworms)
    if err != nil {
        return nil, err
    }
	return bookworms, nil
}
