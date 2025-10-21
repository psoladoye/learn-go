package main

import (
    "flag"
    "fmt"
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

    fmt.Println("Here are the books in common:")
    displayBooks(commonBooks)
}