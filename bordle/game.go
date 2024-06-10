package bordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Length of the word in the game
const solutionLength = 5

// / Type tp represent the game state.
type Game struct {
	word        string
	reader      *bufio.Reader
	maxAttempts int
}

// Function to return a new bordle game.
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		word:        solution,
		maxAttempts: maxAttempts,
	}

	// // Initialise the game
	// g.word = pickRandomWord()

	return g
}

func (g *Game) Play() {
	fmt.Println("Welcome to Bordle!")
	fmt.Printf("Enter a guess\n")
	guess := g.ask()
	fmt.Printf("Your guess is: %s\n", string(guess))
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

		// string(playerInput) wraps the []byte from the readline func
		guess := splitToUppercaseCharacters(string(playerInput))

		// validate the player's guess
		err = g.validateGuess(guess, len(g.word))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s\n", err.Error())
		} else {
			return guess
		}
	}

}

// splitToUppercaseCharacters is a naive implementation to turn a string into a list of characters.
func splitToUppercaseCharacters(input string) []rune {
	// will convert each byte (character) to uppercase
	// +
	// converts to runes
	return []rune(strings.ToUpper(input))
}

// errInvalidWordLength is returned when the guess has the wrong number of characters.
var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

// ensures the guess is at least a valid guess before proceeding
func (g *Game) validateGuess(guess []rune, wordLength int) error {
	if len(guess) != wordLength {
		// _, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution! Expected %d characters, got %d.\n", solutionLength, len(guess))
		return fmt.Errorf("expected %d characters, got %d, %w", solutionLength, len(guess), errInvalidWordLength)
	}
	return nil
}

func pickRandomWord() string {
	return "house"
}
