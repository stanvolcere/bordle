package bordle

import "strings"

// hint describes the validity of a character in a word.
type hint byte

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

// feedback is a list of hints, one per character of the word.
type feedback []hint

// String implements the Stringer interface.
// (Stringer is essentially Go's way of having a
// toString() on an arbirary type)
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "⬜️" // grey square
	case wrongPosition:
		return "🟡" // yellow circle
	case correctPosition:
		return "💚" // green heart
	default:
		// This should never happen.
		return "💔" // red broken heart
	}
}

// // THE BELOW CODE IS WRONG AND IS ONLY PRESENT
// // FOR LEARNING PURPOSES
// // In Go, strings are immutable.
// // Constant. We cannot alter them.
// func (fb feedback) StringConcat() string {
// 	var output string
// 	for _, h := range fb {
// 		output += h.String()
// 	}
// 	return output
// }

func (fb feedback) StringConcat() string {
	// initialse the string builder type
	sb := strings.Builder{}

	for _, h := range fb {
		sb.WriteString(h.String())
	}

	return sb.String()
}
