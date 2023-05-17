package wordle

import "testing"

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
