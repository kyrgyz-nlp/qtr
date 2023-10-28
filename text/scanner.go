package text

type Scanner struct {
	input           []rune
	ch              rune
	prevChar        rune
	currentPosition int
}

func (s *Scanner) readNextChar() {
	if s.currentPosition >= len(s.input) {
		s.ch = -1
	} else {
		s.ch = s.input[s.currentPosition]
		if s.currentPosition > 0 {
			s.prevChar = s.input[s.currentPosition-1]
		}
		s.currentPosition += 1
	}
}

func (s *Scanner) peekNextChar() rune {
	if s.currentPosition >= len(s.input) {
		return -1
	} else {
		return s.input[s.currentPosition]
	}
}

func (s *Scanner) newInput(input string) {
	s.input = []rune(input)
	s.ch = -1
	s.prevChar = -1
	s.currentPosition = 0
}

func NewScanner(input string) *Scanner {
	return &Scanner{
		input:           []rune(input),
		ch:              -1,
		prevChar:        -1,
		currentPosition: 0,
	}
}
