package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Example_main() {
	main()
	// Output:
	// Hello, World!
}

func TestGreet(t *testing.T) {
	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase {
		"English": {
			lang: "en",
			want: "Hello, World!",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Unknown": {
			lang: "un",
			want: "",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t * testing.T) {
			// Act
			got := greet(tc.lang)

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}