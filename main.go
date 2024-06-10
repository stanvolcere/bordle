package main

import (
	"bordle/bordle"
	"os"
)

func main() {
	b := bordle.New(os.Stdin, "house", 3)
	b.Play()

	// it takes 3 bytes to represent non-latin characters
	// fmt.Println(len("Hello, 世界"))
	// fmt.Println(len([]rune("Hello, 世界")))
}
