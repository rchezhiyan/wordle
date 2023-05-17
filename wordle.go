package wordle

import (
	"strings"
)

const (
	maxGuesses = 6
	wordSize   = 5
)

// letterStatus can be none, correct, present, or absent
type letterStatus int

const (
	// none = no status, not guessed yet
	none letterStatus = iota
	// absent = not in the word
	absent
	// present = in the word, but not in the correct position
	present
	// correct = in the correct position
	correct
)

type wordleState struct {
	// word is the word that the user is trying to guess
	word [wordSize]byte
	// guesses holds the guesses that the user has made
	guesses [maxGuesses]guess
	// currGuess is the index of the available slot in guesses
	currGuess int
}

// guess is an attempt to guess the word
type guess [wordSize]letter

type letter struct {
	// char is the letter that this struct represents
	char byte
	// status is the state of the letter (absent, present, correct)
	status letterStatus
}

func newWordleState(word string) wordleState {

	w := wordleState{}
	copy(w.word[:], word)
	return w
}

func newLetter(b byte) letter {
	lt := letter{
		char:   b,
		status: 0,
	}

	return lt

}

func newGuess(gw string) guess {
	g := guess{}

	for i, c := range gw {
		g[i] = newLetter(byte(c))
	}

	return g
}

func (g *guess) string() string {
	str := ""
	for _, l := range g {
		if 'A' <= l.char && l.char <= 'Z' {
			str += string(l.char)
		}
	}
	return str
}

// updateLettersWithWord updates the status of the letters in the guess based on a word
func (g *guess) updateLettersWithWord(word [wordSize]byte) {

	for i := range g {
		l := &g[i]
		if l.char == word[i] {
			l.status = correct
		} else if strings.Contains(string(word[:]), string(l.char)) {
			l.status = present
		} else {
			l.status = absent
		}
	}
}

// appendGuess adds a guess to the wordleState. It returns an error
// if the guess is invalid.
func (w *wordleState) appendGuess(g guess) error {

	guesses := &w.guesses

	guesses.append(g)

}
