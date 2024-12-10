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
	attempt := 1

	for {
		fmt.Printf("Enter a %d-character guess:\n", solutionLength)
		// The ReadLine method will give us the
		// user’s input as a slice of bytes.
		playerInput, _, err := g.reader.ReadLine()

		if err != nil {
			fmt.Fprintf(os.Stderr, "We failed to read your guess %s\n", err.Error())
			continue
		}

		// string(playerInput) wraps the []byte from the readline func
		// returns the chars in the form of []runes
		// aka []int32
		guess := splitToUppercaseCharacters(string(playerInput))

		// validate the player's guess
		err = g.validateGuess(guess, len(g.word))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s\n", err.Error())
			continue
		}

		// check whether the guess is correct or not
		result, success := g.testGuess(guess)

		println("success")
		println(success)

		if attempt == g.maxAttempts || success == 1 {
			return result
		}

		fmt.Printf("Your guess so far: %s\n", string(result))
		fmt.Printf("Attempts remaining: %d\n", g.maxAttempts-attempt)
		fmt.Printf("----------\n")
		attempt++
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

func (g *Game) testGuess(guess []rune) ([]rune, int) {
	// if cound then return a 1 else return 0
	newGuess := make([]rune, len(g.word))

	wordAsRune := splitToUppercaseCharacters(g.word)

	// fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", wordAsRune, wordAsRune, wordAsRune)

	for i, r := range wordAsRune {
		if r == guess[i] {
			newGuess[i] = r
		} else {
			newGuess[i] = 0
		}
	}

	return newGuess, 0
}

func pickRandomWord() string {
	return "house"
}
