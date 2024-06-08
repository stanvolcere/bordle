package bordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Length of the word in the game
const solutionLength = 5

// / Type tp represent the game state.
type Game struct {
	word   string
	reader *bufio.Reader
}

// Function to return a new bordle game.
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}

	// Initialise the game
	g.word = pickRandomWord()

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Bordle!")
	fmt.Printf("Enter a guess\n")
	g.ask()
}

// Function will ask the player for their next guess
// and return slice of runes[]
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", solutionLength)

	for {
		// The ReadLine method will give us the
		// user’s input as a slice of bytes.
		playerInput, _, err := g.reader.ReadLine()

		if err != nil {
			fmt.Fprintf(os.Stderr, "We failed to read your guess %s\n", err.Error())
			continue
		}

		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d.\n", solutionLength, len(guess))
		} else {
			return guess
		}

	}

}

func pickRandomWord() string {
	return "house"
}
