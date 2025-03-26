package futf

func ToUTF8(str string) string {
	raw := []byte(str)
	enc, prob := AllInOne.Guese(raw)

	if enc == nil || prob == 0 {
		return str
	}

	return string(enc.Decode(raw))
}
