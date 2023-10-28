package text

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

type Converter struct {
	scanner      *Scanner
	inputScanner *bufio.Scanner
}

func (c *Converter) Convert(writer io.Writer) {
	matcher := NewMatcher()
	var result []rune
	for c.inputScanner.Scan() {
		result = []rune{}
		line := c.inputScanner.Text()
		c.scanner.newInput(line)
		for {
			c.scanner.readNextChar()
			if c.scanner.ch >= 0 {
				matchResult := matcher.match(
					c.scanner.ch,
					c.scanner.prevChar,
					c.scanner.peekNextChar(),
				)
				for _, char := range matchResult {
					result = append(result, char)
				}
			} else {
				break
			}
		}

		_, err := fmt.Fprint(writer, string(result))
		if err != nil {
			log.Fatal("Natyjcany cazuu uçurunda katalar boldu", err)
		}
	}

	if err := c.inputScanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func NewConverter(scanner *bufio.Scanner) *Converter {
	return &Converter{
		scanner:      NewScanner(""),
		inputScanner: scanner,
	}
}
