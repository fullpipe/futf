package futf

type Guesser interface {
	Guese(str []byte) (Decoder, int)
}
