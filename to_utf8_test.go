package futf

import (
	"testing"
)

func TestToUTF8(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			"Нормальный текст",
			"Нормальный текст",
			"Нормальный текст",
		},
		{
			"w1251 encoded as utf-8",
			"Ìèðîâàÿ ôàíñòàñòèêà îò À äî ß",
			"Мировая фанстастика от А до Я",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUTF8(tt.str); got != tt.want {
				t.Errorf("ToUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}
