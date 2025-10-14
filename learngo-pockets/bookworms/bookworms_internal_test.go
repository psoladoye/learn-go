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
