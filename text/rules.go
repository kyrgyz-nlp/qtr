package text

import "unicode"

type AlphabetMapping map[rune]string

var Mapping = AlphabetMapping{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'ғ': "ğ",
	'д': "d",
	'е': "e",
	'ё': "jo",
	'ж': "c",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'қ': "q",
	'л': "l",
	'м': "m",
	'н': "h",
	'ң': "ñ",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ү': "ü",
	'ф': "f",
	'х': "h",
	'ц': "",
	'ч': "ç",
	'ш': "ş",
	'щ': "ş",
	'ь': "",
	'ъ': "",
	'ы': "y",
	'э': "ä",
	'ә': "ә",
	'ю': "ju",
	'я': "ja",
	'А': "A",
	'Б': "B",
	'В': "V",
	'Г': "G",
	'Ғ': "Ğ",
	'Д': "D",
	'Е': "E",
	'Ё': "JO",
	'Ж': "C",
	'З': "Z",
	'И': "I",
	'Й': "J",
	'К': "K",
	'Қ': "Q",
	'Л': "L",
	'М': "M",
	'Н': "H",
	'Ң': "Ñ",
	'О': "O",
	'П': "P",
	'Р': "R",
	'С': "S",
	'Т': "T",
	'У': "U",
	'Ү': "Ü",
	'Ф': "F",
	'Х': "H",
	'Ц': "",
	'Ч': "Ç",
	'Ш': "Ş",
	'Щ': "Ş",
	'Ь': "",
	'Ъ': "",
	'Ы': "Y",
	'Э': "Ä",
	'Ә': "Ә",
	'Ю': "JU",
	'Я': "JA",
}

type void struct{}
type SpecialCase struct {
	Letter      string
	WithLetters map[rune]void
}

// Özgöço qeisteri q menen ğ
// Oñ ce sol cagynda bolso
// Q(қ),Ğ(ғ) > "а,ы,о,у" coon ündüülör menen cazylat.
// Bizge coon ündüülör erecesin ele dalildöö kerek
// Içke ündüülörgö negizki algoritmge tuuraluu işteit
// K(к),G(г) > "ә,е,и,ө,ү" içke ündüülör menen cazylat.
var withLetters = map[rune]void{
	'а': {}, 'ы': {}, 'о': {}, 'у': {},
	'А': {}, 'Ы': {}, 'О': {}, 'У': {},
}

var specialCases = map[rune]SpecialCase{
	'к': {
		Letter:      "q",
		WithLetters: withLetters,
	},
	'К': {
		Letter:      "Q",
		WithLetters: withLetters,
	},
	'г': {
		Letter:      "ğ",
		WithLetters: withLetters,
	},
	'Г': {
		Letter:      "Ğ",
		WithLetters: withLetters,
	},
}

type Alphabet struct {
	Mapping      AlphabetMapping
	SpecialCases map[rune]SpecialCase
}

type Matcher interface {
	match(char rune, prevChar, nextChar rune) []rune
}

type RuleMatcher struct {
	alphabet Alphabet
}

func (r *RuleMatcher) match(char rune, prevChar, nextChar rune) []rune {
	if unicode.IsSpace(char) {
		return []rune{char}
	}

	// Özgöço qeisterğe dal kelse
	if specialCase, found := r.alphabet.SpecialCases[char]; found {
		_, prevCharMatched := specialCase.WithLetters[prevChar]
		_, nextCharMatched := specialCase.WithLetters[nextChar]
		if prevCharMatched || nextCharMatched {
			return []rune(specialCase.Letter)
		}
	}

	// Calpy qalyby
	if value, found := r.alphabet.Mapping[char]; found {
		return []rune(value)
	} else {
		return []rune{char}
	}
}

func NewMatcher() Matcher {
	return &RuleMatcher{
		alphabet: Alphabet{
			Mapping:      Mapping,
			SpecialCases: specialCases,
		},
	}
}
