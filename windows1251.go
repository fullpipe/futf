package futf

import (
	"bytes"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Windows1251 struct{}

var missCodeWindows1251Cyrillic = map[string]bool{}

func init() {
	var rc rune = 0x0400
	for rc < 0x04FF {
		missCoded := Dec(Enc(string(rc), charmap.Windows1251), charmap.Windows1252)
		if len(missCoded) == 0 {
			rc += 1
			continue
		}

		missCodeWindows1251Cyrillic[string(missCoded)] = true
		rc += 1
	}
}

func Dec(s []byte, enc encoding.Encoding) []byte {
	dec := enc.NewDecoder()
	out, _ := dec.Bytes(s)
	return out
}

func Enc(s string, enc encoding.Encoding) []byte {
	var b bytes.Buffer
	wr := transform.NewWriter(&b, enc.NewEncoder())
	wr.Write([]byte(s))
	wr.Close()

	return b.Bytes()
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func (g *Windows1251) Guese(str []byte) (Decoder, int) {
	maxLen := 0
	total := 0
	words := 0
	l := 0

	for _, b := range string(str) {
		if missCodeWindows1251Cyrillic[string(b)] {
			total += 1
			l += 1
		} else {
			if l > 4 {
				words += 1
			}
			maxLen, l = max(l, maxLen), 0
		}
	}
	if l > 3 {
		words += 1
	}
	maxLen, l = max(l, maxLen), 0

	decoder := &MissCodedDecoder{origin: charmap.Windows1251, misscoded: charmap.Windows1252}

	if words > 1 {
		return decoder, 100
	}

	if maxLen > 4 {
		return decoder, 75
	}

	if maxLen > 3 || total > len(str)/2 {
		return decoder, 75
	}

	if maxLen > 2 {
		return decoder, 50
	}

	if maxLen > 1 {
		return decoder, 25
	}

	return decoder, 0
}
