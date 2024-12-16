package main

import (
	"bordle/bordle"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// Get dictionary source from env file
	dictionary_path := os.Getenv("DICTIONARY_SOURCE")

	// Check if the variable is set
	if dictionary_path == "" {
		fmt.Printf("%s is not set or empty\n", dictionary_path)
	} else {
		fmt.Printf("%s=%s\n", "env", dictionary_path)
	}

	// Load the dictionary into game state
	dictionaryWords, err := bordle.LoadDictionary(dictionary_path)
	if err != nil {
		fmt.Println("Error loading dictionary:", err)
		return
	}
	dictionary := bordle.NewDictionary(dictionaryWords)

	rounds := 6

	b := bordle.New(os.Stdin,
		bordle.PickRandomWord(),
		rounds,
		dictionary,
	)
	b.Play()

	// it takes 3 bytes to represent non-latin characters
	// fmt.Println(len("Hello, 世界"))
	// fmt.Println(len([]rune("Hello, 世界")))
}
