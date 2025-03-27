package futf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
)

// TODO: they shoould pass. Its required to make Windows1251 guesser more strict
func TestWindows1251_Guese(t *testing.T) {
	// t.Log(string(toEnc("Ñàíêò-Ïåòåðáóðã", charmap.ISO8859_1)))
	// t.Log(string(dec(toEnc("Ñàíêò-Ïåòåðáóðã", charmap.ISO8859_1), charmap.Windows1251)))
	// t.Log(string(Dec(Enc("Ñàíêò-Ïåòåðáóðã", charmap.Windows1252), charmap.Windows1251)))

	// t.Log(string(dec([]byte("Ñàíêò-Ïåòåðáóðã"), charmap.ISO8859_1)))
	// t.Log(string(dec([]byte("Ñàíêò-Ïåòåðáóðã"), charmap.ISO8859_1)))
	// t.Log(string(dec([]byte("Ñàíêò-Ïåòåðáóðã"), charmap.Windows1251)))
	// t.Log(string(dec(dec([]byte("Ñàíêò-Ïåòåðáóðã"), charmap.ISO8859_1), charmap.Windows1251)))
	// t.Log(string(dec(dec([]byte("Ñàíêò-Ïåòåðáóðã"), charmap.ISO8859_1), charmap.Windows1252)))
	tests := []struct {
		name     string
		str      []byte
		wantProb int
	}{
		{
			"long cyrillic sequence",
			Enc("Привет!", charmap.Windows1251),
			0,
		},
		{
			"cyrillic chars separated",
			Enc("П-р_и%в*е,т", charmap.Windows1251),
			0,
		},
		{
			"short cyrillic sequence",
			Enc("При!!!!", charmap.Windows1251),
			0,
		},
		{
			"super short cyrillic sequence",
			Enc("Пр!!!!", charmap.Windows1251),
			0,
		},
		{
			"Санкт-Петербург", //
			[]byte("Ð¡Ð°Ð½ÐºÑ‚-ÐŸÐµÑ‚ÐµÑ€Ð±ÑƒÑ€Ð³"),
			0,
		},
		{
			"Санкт-Петербург", //
			[]byte("Ñàíêò-Ïåòåðáóðã"),
			100,
		},
	}

	// t.Log(string(dec("Ñàíêò-Ïåòåðáóðã", charmap.Windows1251)))
	// t.Log(string(dec("Ñàíêò-Ïåòåðáóðã", charmap.Windows1252)))
	// t.Log(string(dec("Ñàíêò-Ïåòåðáóðã", charmap.ISO8859_1)))

	// t.Log(string(Enc("Санкт-Петербург", charmap.Windows1251)))
	// t.Log(string(toEnc("Санкт-Петербург", charmap.Windows1252)))
	// t.Log(string(toEnc("Ñàíêò-Ïåòåðáóðã", charmap.Windows1251)))
	// t.Log(string(toEnc("Ñàíêò-Ïåòåðáóðã", charmap.Windows1252)))
	// t.Log(string(toEnc("Ð¡Ð°Ð½ÐºÑ‚-ÐŸÐµÑ‚ÐµÑ€Ð±ÑƒÑ€Ð³", charmap.Windows1251)))
	// t.Log(string(toEnc("Ð¡Ð°Ð½ÐºÑ‚-ÐŸÐµÑ‚ÐµÑ€Ð±ÑƒÑ€Ð³", charmap.Windows1252)))
	// t.Log(string(toEnc("Санкт-Петербург", charmap.Windows1252)))
	// t.Log(string(toEnc("Ð¡Ð°Ð½ÐºÑ‚-ÐŸÐµÑ‚ÐµÑ€Ð±ÑƒÑ€Ð³", charmap.UtfBOM)))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Windows1251{}
			_, prob := g.Guese(tt.str)
			assert.Equal(t, tt.wantProb, prob)
		})
	}
}
