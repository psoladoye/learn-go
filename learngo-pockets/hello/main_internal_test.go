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

func TestGreet_English(t *testing.T) {
	// Arrange
	want := "Hello, World!"
	lang := language("en")

	// Act
	got := greet(lang)

	// Assert
	assert.Equal(t, got, want)
}

func TestGreet_French(t *testing.T) {
	// Arrage
	lang := language("fr")
	want := "Bonjour le monde"

	// Act
	got := greet(lang)

	// Assert
	assert.Equal(t, want, got)

}

func TestGreet_Unknown(t *testing.T) {
	// Arrange
	lang := language("un")
	want := ""

	// Act
	got := greet(lang)

	// Assert
	assert.Equal(t, want, got)
}
