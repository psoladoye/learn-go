package main

import (
    "flag"
    "fmt"
    "sort"
)

func main() {
    var filepath string
    flag.StringVar(&filepath, "filepath", "testdata/bookworms.json", "file path to testdata/bookworms.json")
    flag.Parse()

    bookworms, err := loadBookworms(filepath)
    if err != nil {
        fmt.Printf("failed to load bookworms testdata %v", err)
        return
    }

    commonBooks := findCommonBooks(bookworms)

    books := []Book{
        { Author: "C", Title: "3" },
        { Author: "A", Title: "1" },
        { Author: "B", Title: "2" },
        { Author: "B", Title: "3" },
    }

    sort.Sort(byAuthor(books))
    displayBooks(books)

    fmt.Println("\nHere are the books in common:")
    displayBooks(commonBooks)
}