package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
				if t.times == 1 {
					commonBooks = append(commonBooks, b)
				}
			} else {
				register[b.Title] = tracker{ book: b, times: 1}
			}
		}
	}
	return sortBooks(commonBooks)
}

// sortBooks sorts the books first by Author then Title.
func sortBooks(books []Book) []Book {
	sort.Sort(byAuthor(books))
	return books
}

// displayBooks prints out the titles and authors of a list of books
func displayBooks(books []Book) {
	for _, b := range books {
		fmt.Printf("- %s by %s \n", b.Title, b.Author)
	}
}

// byAuthor is a list of Book.
// Definig a custom type to implement the interface
type byAuthor []Book

// Len implements the sort.Interface by returning the length of the BookByAuthor
func (b byAuthor) Len() int { return len(b) }

// Swap implements the sort.Interface and swaps two books.
func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// Less implements sort.Interface and returns
// books sorted by Author and Title.
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}
