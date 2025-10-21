package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	handmaidsTale = Book{
		Author: "Margaret Atwood", Title: "The Handmaid's Tale",
	}
	oryxAndCrake = Book{
		Author: "Margaret Atwood", Title: "Oryx and Crake",
	}
    theBellJar = Book{
        Author: "Sylvia Plath", Title: "The Bell Jar",
    }
    janeEyre = Book{
        Author: "Charlotte Brontë", Title: "Jane Eyre",
    }
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		name     string
		filepath string
		want     []Bookworm
		wantErr  bool
	}

	tests := []testCase{
		{
			name:     "File does not exist",
			filepath: "testdata/none.json",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "File exists",
			filepath: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
	}

    for _, tc := range tests {
        t.Run(tc.name, func(t * testing.T) {
            // Act
            got, err := loadBookworms(tc.filepath)

			// Assert
            if tc.wantErr {
				assert.Nil(t, got)
                require.Error(t, err)
				return
            }
			assert.ElementsMatch(t, tc.want, got)
        })
    }
}

func TestFindCommonBooks(t *testing.T) {
	// Arrange
	type testCase struct {
		scenario string
		bookworms []Bookworm
		want []Book
	}

	tests := []testCase{
		{
			scenario: "nobody has any books",
			bookworms: []Bookworm{
				{ Name: "Author A", Books: []Book{} },
				{ Name: "Author B", Books: []Book{} },
				{ Name: "Author C", Books: []Book{} },
			},
			want: []Book{},
		},
		{
			scenario: "one bookworm has no books",
			bookworms: []Bookworm{
				{ Name: "Author A", Books: []Book{handmaidsTale, theBellJar} },
				{ Name: "Author B", Books: []Book{} },
			},
			want: []Book{},
		},
		{
			scenario: "more than 2 bookworms have a book in common",
			bookworms: []Bookworm{
				{ Name: "Author A", Books: []Book{handmaidsTale, theBellJar} },
				{ Name: "Author B", Books: []Book{} },
				{ Name: "Author C", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre} },
			},
			want: []Book{handmaidsTale},
		},
		{
			scenario: "people have no books in common",
			bookworms: []Bookworm{
				{ Name: "Author A", Books: []Book{handmaidsTale, theBellJar} },
				{ Name: "Author B", Books: []Book{oryxAndCrake} },
				{ Name: "Author C", Books: []Book{janeEyre} },
			},
			want: []Book{},
		},
		{
			scenario: "everyone has read the same books",
			bookworms: []Bookworm{
				{ Name: "Author A", Books: []Book{handmaidsTale, theBellJar} },
				{ Name: "Author B", Books: []Book{handmaidsTale, theBellJar} },
				{ Name: "Author C", Books: []Book{handmaidsTale, theBellJar} },
			},
			want: []Book{handmaidsTale, theBellJar},
		},
	}

	for _, tc := range tests {
		t.Run(tc.scenario, func(t *testing.T) {
			// Act
			got := findCommonBooks(tc.bookworms)

			// Assert
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
