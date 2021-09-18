package tokens

import "errors"

var InvalidParenthesesError = errors.New("invalid parentheses")

func Tokenize(expr string) Tokenizer {
    return Tokenizer{
        expr: expr,
    }
}

type Tokenizer struct {
    expr string
    lastPos int
    err error
    curr Token
    parenthesesCnt int
}

func (tk *Tokenizer) Next() bool {
    if tk.lastPos >= len(tk.expr) || tk.err != nil {
        return false
    }

    s := tk.expr[tk.lastPos]
    rawToken := ""
    if s >= '0' && s <= '9' {
        rawToken = tk.readNumber(tk.expr, tk.lastPos)
    } else {
        rawToken = string(s)
    }

    tok, err := NewToken(rawToken)
    if err != nil {
        tk.err = err
        return false
    }

    err = tk.accountParentheses(tok)
    if err != nil {
        tk.err = err
        return false
    }

    tk.curr = tok
    tk.lastPos = tk.lastPos + len(rawToken)


    return true
}

func (tk *Tokenizer) accountParentheses(tok Token) error {
    if tok.Type() == OpenParenthesisToken {
        tk.parenthesesCnt++
    } else if tok.Type() == CloseParenthesisToken {
        tk.parenthesesCnt--
    }

    if tk.parenthesesCnt < 0 {
        return InvalidParenthesesError
    }

    if tk.lastPos + len(tok.String()) >= len(tk.expr) && tk.parenthesesCnt > 0 {
        return InvalidParenthesesError
    }

    return nil
}

func (tk *Tokenizer) Curr() Token {
    return tk.curr
}

func (tk *Tokenizer) Err() error {
    return tk.err
}

func (tk *Tokenizer) readNumber(expr string, from int) string {
    n := ""
    for i := from; i < len(expr); i++ {
        s := expr[i]
        if s >= '0' && s <= '9' {
            n += string(s)
        } else {
            return n
        }
    }

    return n
}
