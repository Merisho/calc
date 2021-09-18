package tokens

import (
    "github.com/stretchr/testify/suite"
    "testing"
)

func TestTokenizer(t *testing.T) {
    suite.Run(t, new(TokenizerTestSuite))
}

type TokenizerTestSuite struct {
    suite.Suite
}

func (ts *TokenizerTestSuite) TestBasicTokenization() {
    tok := Tokenize("321412341+9999999999")

    ok := tok.Next()
    ts.NoError(tok.Err())
    ts.True(ok)

    t := tok.Curr()
    ts.Equal(NumberToken, t.Type())
    num, _ := t.Number()
    ts.EqualValues(321412341, num)

    ok = tok.Next()
    ts.NoError(tok.Err())
    ts.True(ok)

    t = tok.Curr()
    ts.Equal(PlusToken, t.Type())
    _, ok = t.Number()
    ts.False(ok)

    ok = tok.Next()
    ts.NoError(tok.Err())
    ts.True(ok)

    t = tok.Curr()
    ts.Equal(NumberToken, t.Type())
    num, _ = t.Number()
    ts.EqualValues(9999999999, num)

    ok = tok.Next()
    ts.NoError(tok.Err())
    ts.False(ok)
}

func (ts *TokenizerTestSuite) TestInvalidToken() {
    tok := Tokenize("1+_5")

    // Move to 1
    tok.Next()
    // Move to +
    tok.Next()

    // Move to _ which is an invalid symbol
    ok := tok.Next()
    ts.False(ok)
    ts.EqualError(tok.Err(), "unknown token")
}

func (ts *TokenizerTestSuite) TestTokenization() {
    expr := "237519+123123-2415+123521-2-(5-6)+2351+(5234-(4231+5234-(543521)))*(5+2*3)"
    expectedTokenTypeSequence := []TokenType{
        NumberToken,
        PlusToken,
        NumberToken,
        MinusToken,
        NumberToken,
        PlusToken,
        NumberToken,
        MinusToken,
        NumberToken,
        MinusToken,
        OpenParenthesisToken,
        NumberToken,
        MinusToken,
        NumberToken,
        CloseParenthesisToken,
        PlusToken,
        NumberToken,
        PlusToken,
        OpenParenthesisToken,
        NumberToken,
        MinusToken,
        OpenParenthesisToken,
        NumberToken,
        PlusToken,
        NumberToken,
        MinusToken,
        OpenParenthesisToken,
        NumberToken,
        CloseParenthesisToken,
        CloseParenthesisToken,
        CloseParenthesisToken,
        MultiplicationToken,
        OpenParenthesisToken,
        NumberToken,
        PlusToken,
        NumberToken,
        MultiplicationToken,
        NumberToken,
        CloseParenthesisToken,
    }

    tok := Tokenize(expr)

    for tok.Next() {
        ts.Require().NoError(tok.Err())
        ts.Equal(expectedTokenTypeSequence[0], tok.Curr().Type())
        expectedTokenTypeSequence = expectedTokenTypeSequence[1:]
    }

    ts.Require().NoError(tok.Err())
}

func (ts *TokenizerTestSuite) TestInvalidParentheses() {
    cases := []string{
        ")-1+2",
        "(4+5))",
        "(3+(5-4)",
        "8+(2-(5+7)))",
        "2+3+5+(",
    }

    for _, c := range cases {
        tok := Tokenize(c)
        for tok.Next() {}
        ts.EqualError(tok.Err(), "invalid parentheses")
    }
}
