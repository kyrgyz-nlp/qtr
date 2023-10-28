package text

import "bytes"

// dropCR aqyrqy \r simvolun alyp salat
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

// ScanLinesWithLF Scanner'ge saptardy bölüü funktsijasy teksttin ar bir sabyn
// qaitaryp cañy sap belgisi menen qoşo beret al boş bolup qalyş mümkün.
// Cañy sap belgisi bul koşumça simvol bolot, eñ aqyrqysy cańy sap belgisi kelet.
// Regularduu tujuntma notatsijasy - `\r?\n`.
// Cańy sap belgisi bolboso da aqyrqy maanilüü saby soñunda qaitarylat.
func ScanLinesWithLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0 : i+1]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
