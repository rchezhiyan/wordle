package wordle

import (
	"testing"

	"github.com/rchezhiyan/wordle/words"
)

func TestNewWordleState(t *testing.T) {

	word := "HELLO"
	ws := newWordleState(word)
	wordleAsString := string(ws.word[:])

	t.Log("Created wordleState:")
	t.Logf("		word: %s", wordleAsString)
	t.Logf("		guesses: %v", ws.guesses)
	t.Logf("		currGuess: %v", ws.currGuess)

	if wordleAsString != word {
		t.Errorf("word = %s; want %s", wordleAsString, word)
	}

}

func statusToString(status letterStatus) string {
	switch status {
	case none:
		return "none"
	case correct:
		return "correct"
	case present:
		return "present"
	case absent:
		return "absent"
	default:
		return "unknown"
	}
}
func TestNewGuess(t *testing.T) {

	guess := "LOVER"

	wg := newGuess(guess)

	t.Logf("Newguess: %v", wg.string())

	for i, l := range wg {

		if l.char != guess[i] || l.status != none {

			t.Errorf(
				"letter[%d] = %c, %s; want %c, none",
				i,
				l.char,
				statusToString(l.status),
				guess[i],
			)
		}

	}

}

func TestUpdateLettersWithWord(t *testing.T) {
	guessWord := "YIELD"
	guess := newGuess(guessWord)

	var word [wordSize]byte
	copy(word[:], "HELLO")
	guess.updateLettersWithWord(word)

	statuses := []letterStatus{
		absent,  // "Y" is not in "HELLO"
		absent,  // "I" is not in "HELLO"
		present, // "E" is in "HELLO" but not in the correct position
		correct, // "L" is in "HELLO" and in the correct position
		absent,  // "D" is not in "HELLO"
	}

	// Check that statuses are correct

	for i, l := range guess {

		if l.status != statuses[i] {
			t.Errorf(
				"letter[%d] = %c, %s; want %c, %s",
				i,
				l.char,
				statusToString(l.status),
				guessWord[i],
				statusToString(statuses[i]),
			)
		}

	}

}

func TestAppendGuess(t *testing.T) {
	ws := newWordleState("HELLO")

	err := ws.appendGuess(newGuess("YIELD"))

	if err == nil {
		if ws.currGuess != 1 {
			t.Errorf("currGuess = %d; wnat 1", ws.currGuess)
		}
	}

	for i := 0; i < 5; i++ {
		guess := newGuess(words.GetWord())
		if err := ws.appendGuess(guess); err != nil {
			t.Errorf("newGuess() returned an error: %v", err)
		}
	}
}

func TestAppendGuessError(t *testing.T) {

	ws := newWordleState("HELLO")

	// var g guess
	// Add six guesses
	for i := 0; i < 6; i++ {
		g := newGuess(words.GetWord())
		ws.appendGuess(g)
	}

	g := newGuess(words.GetWord())
	if err := ws.appendGuess(g); err == nil {
		t.Errorf("appendGuess sholud result in an error")
	}

}

func TestIsWordGuessed(t *testing.T) {
	ws := newWordleState("HELLO")
	g := newGuess("HELLO")

	g.updateLettersWithWord(ws.word)
	ws.appendGuess(g)

	if !ws.isWordGuessed() {
		t.Errorf("isWordGuessed should return true")
	}
}

func TestShouldEndGameCorrect(t *testing.T) {
	ws := newWordleState("HELLO")
	g := newGuess("HELLO")

	g.updateLettersWithWord(ws.word)
	ws.appendGuess(g)

	if !ws.shouldEndGame() {
		t.Errorf("shouldEndGame should return true")
	}

}

func TestShouldEndGameMaxGuesses(t *testing.T) {
	ws := newWordleState("HELLO")

	// add 6 wrong guesses
	for i := 0; i < 6; i++ {
		g := newGuess(words.GetWord())
		g.updateLettersWithWord(ws.word)
		ws.appendGuess(g)
	}
	if !ws.shouldEndGame() {
		t.Errorf("shouldEndGame should return true")
	}
}
