package bordle

import "fmt"

// / Type tp represent the game state.
type Game struct {
	word string
}

// Function to return a new bordle game.
func New() *Game {
	g := &Game{}

	// Initialise the game
	g.word = pickRandomWord()

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Bordle!")
	fmt.Printf("Enter a guess\n")
}

func pickRandomWord() string {
	return "house"
}
