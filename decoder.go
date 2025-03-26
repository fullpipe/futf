package futf

import (
	"bytes"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

type Decoder interface {
	Decode(str []byte) []byte
}

type CharmapDecoder struct {
	encoding.Encoding
}

func (d *CharmapDecoder) Decode(str []byte) []byte {
	var b bytes.Buffer
	wr := transform.NewWriter(&b, d.NewDecoder())
	_, err := wr.Write(str)
	if err != nil {
		return str
	}
	return b.Bytes()
}

// Solves cases when original text in Windows1251 was missinterpreted as Windows1252 and than was encoded to unicode
type MissCodedDecoder struct {
	origin    encoding.Encoding
	misscoded encoding.Encoding
}

func (d *MissCodedDecoder) Decode(str []byte) []byte {
	return Dec(Enc(string(str), d.misscoded), d.origin)
}
