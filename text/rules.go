package text

var mapping = map[string]string{
	"а": "a",
	"б": "b",
	"в": "v",
	"г": "g",
	"ғ": "ğ",
	"д": "d",
	"е": "e",
	"ё": "jo",
	"ж": "c",
	"з": "z",
	"и": "i",
	"й": "j",
	"к": "g",
	"қ": "q",
	"л": "l",
	"м": "m",
	"н": "h",
	"ң": "ñ",
	"о": "o",
	"п": "p",
	"р": "r",
	"с": "s",
	"т": "t",
	"у": "u",
	"ү": "ü",
	"ф": "f",
	"х": "h",
	"ц": "",
	"ч": "ç",
	"ш": "ş",
	"щ": "ş",
	"ъ": "",
	"ы": "y",
	"ь": "",
	"э": "ä",
	"ю": "ju",
	"я": "ja",
}

type Side string

const (
	Left  = "left"
	Right = "right"
	All   = "all"
)

type SpecialCase struct {
	WhenPosition Side
	WhenLetter   string
	WithLetters  string
}

// Özgöço qeisteri q menen ğ
// Oñ ce sol cagynda bolso
// Q(қ),Ğ(ғ) > "а,ы,о,у" coon ündüülör menen cazylat.
// Içke ündüülörgö bul negizki algoritmi tuura işteit ⬇️
// K(к),G(г) > "ә,е,и,ө,ү" içke ündüülör menen cazylat.
var specialCase = map[string]*SpecialCase{
	"к": {
		WhenLetter:   "ğ",
		WhenPosition: All,
		WithLetters:  "аыоу",
	},
	"г": {
		WhenLetter:   "ğ",
		WhenPosition: All,
		WithLetters:  "аыоу",
	},
}
