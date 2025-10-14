package main

import (
    "flag"
    "fmt"
)

func main() {
    var filepath string
    flag.StringVar(&filepath, "filepath", "testdata/bookworms.json", "file path to testdata/bookworms.json")
    flag.Parse()

    bookworms, err := loadBookworms("testdata/bookworms.json")
    if err != nil {
        fmt.Printf("failed to load bookworms testdata %v", err)
        return
    }

    commonBooks := findCommonBooks(bookworms)

    for i := range commonBooks {
        fmt.Printf("index=%d, book.name=%s\n", i,commonBooks[i].Title)
    }
}