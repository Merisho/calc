package tokens

type Parenthesis bool

func (p Parenthesis) Open() bool {
    return p == true
}
