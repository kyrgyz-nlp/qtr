package text

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConverter(t *testing.T) {
	// Q(қ),Ğ(ғ) > "а,ы,о,у" coon ündüülör menen cazylat.
	testCases := [][]string{
		{"Кыргыз", "Qyrğyz"},
		{"Казак", "Qazaq"},
		{"Бука", "Buqa"},
		{"Булак", "Bulaq"},
		{"Айгыр", "Ajğyr"},
		{"Тагдыр", "Tağdyr"},
		{"Сыргак", "Syrğaq"},
		{"Максат", "Maqsat"},
		{"Саламаттык сактоо кассасы", "Salamattyq saqtoo qassasy"},
		{
			`# Qytyr qytyr
Саламаттык сактоо кассасы
`,
			`# Qytyr qytyr
Salamattyq saqtoo qassasy
`,
		},
	}

	converter := NewConverter(nil)
	writeBuf := new(bytes.Buffer)
	for _, testCase := range testCases {
		given, expected := testCase[0], testCase[1]
		readerBuf := strings.NewReader(given)
		scanner := bufio.NewScanner(readerBuf)
		scanner.Split(ScanLinesWithLF)
		converter.inputScanner = scanner

		converter.Convert(writeBuf)
		assert.Equal(t, expected, writeBuf.String())
		writeBuf.Reset()
	}
}
