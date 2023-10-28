package text

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScannerBasicFunctions(t *testing.T) {
	testText := "Миңдеген адамдын арасында сен жалгызсың, жалгыз өзүң калганда да жалгызсың."
	var chars string
	scanner := NewScanner(testText)
	assert.Equal(t, 0, scanner.currentPosition)
	scanner.readNextChar()
	chars += string(scanner.ch)
	assert.Equal(t, string(scanner.ch), "М")
	assert.Equal(t, scanner.currentPosition, 1)
	nextChar := scanner.peekNextChar()
	assert.Equal(t, string(nextChar), "и")
	assert.Equal(t, scanner.currentPosition, 1)
	for {
		scanner.readNextChar()
		if scanner.ch == -1 {
			break
		}
		chars += string(scanner.ch)
	}

	assert.Equal(t, testText, chars)
}
