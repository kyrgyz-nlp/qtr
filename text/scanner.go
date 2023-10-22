package text

type Scanner struct {
	input           []rune
	currentPosition int
	ch              rune
}

func (s *Scanner) readNextChar() {
	if s.currentPosition >= len(s.input) {
		s.ch = -1
	} else {
		s.ch = s.input[s.currentPosition]
	}
	s.currentPosition += 1
}

func (s *Scanner) peekNextChar() rune {
	if s.currentPosition >= len(s.input) {
		return 0
	} else {
		return s.input[s.currentPosition]
	}
}

func NewScanner(input string) *Scanner {
	return &Scanner{
		input:           []rune(input),
		currentPosition: 0,
		ch:              0,
	}
}
