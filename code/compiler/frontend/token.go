package frontend

type Token struct {
	Tag int
}
type Num struct {
	Token
	Value int
}
type Word struct {
	Token
	Lexeme string
}

func NewToken(t int) *Token {
	return &Token{t}
}
func NewNum(t int) *Num {
	return &Num{Token{t}, t}
}
func NewWord(t int, l string) *Word {
	return &Word{Token{t}, l}
}
