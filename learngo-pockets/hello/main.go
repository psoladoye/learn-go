package main

import "fmt"

type language string

var phrasebook = map[language]string {
	"en": "Hello, World!",
	"fr": "Bonjour le monde",
	"el": "Χαίρετε Κόσμε",     // Greek
	"he": "שלום עולם",         // Hebrew
    "ur": "ہیلو دنیا",         // Urdu
    "vi": "Xin chào Thế Giới", // Vietnamese
}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return ""
	}
	return greeting
}

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}
