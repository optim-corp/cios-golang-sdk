package wrapper

func (c Char) AsString() string {
	return string(c)
}
func (c Char) AsRune() rune {
	return rune(c)
}
func (c Char) IsAlphabet() bool {
	return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}
func (c Char) IsNumeric() bool {
	return MakeString(c.AsString()).IsNumeric()
}
