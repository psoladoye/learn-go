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
	
	if err = json.NewDecoder(f).Decode(&bookworms); err != nil {
		return nil, err
	}
	return bookworms, nil
}

// findCommonBooks returns books that are on more than one bookworms shelf.
func findCommonBooks(bookworms []Bookworm) []Book {
	type tracker struct {
		book  Book
		times int
	}

	register := map[string]tracker{}
	commonBooks := []Book{}

	for _, bw := range bookworms {
		for _, b := range bw.Books {
			if t, exists := register[b.Title]; exists {
				register[b.Title] = tracker{ book: b, times: t.times + 1}
				commonBooks = append(commonBooks, b)
			} else {
				register[b.Title] = tracker{ book: b, times: 1}
			}
		}
	}
	return commonBooks
}
