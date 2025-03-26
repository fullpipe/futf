package futf

var AllInOne = allInOneGuesser{
	guessers: []Guesser{
		&UtfBOM{},
		&Windows1251{},
	},
}

type allInOneGuesser struct {
	guessers []Guesser
}

func (g *allInOneGuesser) Guese(str []byte) (Decoder, int) {
	var enc Decoder
	prob := 0

	for _, g := range g.guessers {
		e, p := g.Guese(str)
		if p > prob {
			prob = p
			enc = e
		}
	}

	return enc, prob
}
