package futf

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

type UtfBOM struct{}

func (g *UtfBOM) Guese(str []byte) (Decoder, int) {
	enc, hasBOMed := detectFromBOM(str[:5])
	if hasBOMed {
		return &CharmapDecoder{Encoding: enc}, 100
	}

	return nil, 0
}

// TODO: https://stackoverflow.com/a/5830708  -> DetectUnicodeInByteSampleByHeuristics
func detectFromBOM(bom []byte) (encoding.Encoding, bool) {
	if len(bom) < 2 {
		return nil, false
	}

	if len(bom) < 4 || bom[2] != 0 || bom[3] != 0 {
		if bom[0] == 0xff && bom[1] == 0xfe {
			return unicode.UTF16(unicode.LittleEndian, unicode.UseBOM), true
		}
	}

	if bom[0] == 0xfe && bom[1] == 0xff {
		return unicode.UTF16(unicode.BigEndian, unicode.UseBOM), true
	}

	if len(bom) < 3 {
		return nil, false
	}

	if bom[0] == 0xef && bom[1] == 0xbb && bom[2] == 0xbf {
		return unicode.UTF8, true
	}

	// TODO: https://github.com/cention-sany/utf7
	// if bom[0] == 0x2b && bom[1] == 0x2f && bom[2] == 0x76 {
	// 	return unicode.UTF7, true
	// }

	if len(bom) < 4 {
		return nil, false
	}

	if bom[0] == 0xff && bom[1] == 0xfe && bom[2] == 0 && bom[3] == 0 {
		return utf32.UTF32(utf32.LittleEndian, utf32.UseBOM), true
	}

	if bom[0] == 0 && bom[1] == 0 && bom[2] == 0xfe && bom[3] == 0xff {
		return utf32.UTF32(utf32.BigEndian, utf32.UseBOM), true
	}

	return nil, false
}
