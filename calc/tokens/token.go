package tokens

import (
    "errors"
    "strconv"
)

type TokenType string

const (
    NumberToken TokenType = "number"
    PlusToken   TokenType = "+"
    MinusToken  TokenType = "-"
    MultiplicationToken   TokenType = "*"
    DivisionToken   TokenType = "/"
    OpenParenthesisToken  TokenType = "("
    CloseParenthesisToken TokenType = ")"
)

func NewToken(rawToken string) (Token, error) {
    num, err := strconv.ParseInt(rawToken, 10, 64)
    if err == nil {
        return Token{
            token: rawToken,
            num:   num,
            tp:    NumberToken,
        }, nil
    }

    switch rawToken {
    case "+":
        return Token{
            token: rawToken,
            tp: PlusToken,
        }, nil
    case "-":
        return Token{
            token: rawToken,
            tp: MinusToken,
        }, nil
    case "(":
        return Token{
            token: rawToken,
            tp: OpenParenthesisToken,
        }, nil
    case ")":
        return Token{
            token: rawToken,
            tp: CloseParenthesisToken,
        }, nil
    case "*":
        return Token{
            token: rawToken,
            tp: MultiplicationToken,
        }, nil
    case "/":
        return Token{
            token: rawToken,
            tp: DivisionToken,
        }, nil
    }

    return Token{}, errors.New("unknown token")
}

type Token struct {
    token string
    num   int64
    tp    TokenType
}

func (t Token) String() string {
    return t.token
}

func (t Token) Number() (int64, bool) {
    if t.tp == NumberToken {
        return t.num, true
    }

    return 0, false
}

func (t Token) Type() TokenType {
    return t.tp
}

func (t Token) Operator() (Operator, bool) {
    if t.isOperator() {
        return Operator(t.token), true
    }

    return "", false
}

func (t Token) isOperator() bool {
    return t.tp == PlusToken || t.tp == MinusToken || t.tp == MultiplicationToken || t.tp == DivisionToken
}

func (t Token) Parenthesis() (Parenthesis, bool) {
    if t.tp == OpenParenthesisToken || t.tp == CloseParenthesisToken {
        return t.tp == OpenParenthesisToken, true
    }

    return false, false
}
