package calc

import (
    "github.com/stretchr/testify/suite"
    "testing"
)

func TestCalc(t *testing.T) {
       suite.Run(t, new(CalcTestSuite))
}

type CalcTestSuite struct {
   suite.Suite
}

func (ts *CalcTestSuite) TestCalc() {
   calc := NewCalc()

   cases := []struct{
       expr string
       expected string
   }{
       {
           "(2 + 2 * 2) / 3",
           "2",
       },
       {
           "2 + 3",
           "5",
       },
       {
           "7 + 5",
           "12",
       },
       {
           "57239 + 124245",
           "181484",
       },
       {
           "-1 + 2",
           "1",
       },
       {
           "+1  +   2",
           "3",
       },
       {
           "10 - 7",
           "3",
       },
       {
           "-5 - 7",
           "-12",
       },
       {
           "2 + (3 - (5 + 2))",
           "-2",
       },
       {
         "4 * 5",
         "20",
       },
       {
           "-10 / 5",
           "-2",
       },
       {
           "-(2 + 3)",
           "-5",
       },
       {
           "24 / 3 / 4",
           "2",
       },
       {
           "1 + 56 / 7",
           "9",
       },
       {
           "11 * 12 / 4",
           "33",
       },
       {
           "(1 + 55) / 7",
           "8",
       },
   }

   for _, c := range cases {
       res, err := calc.Calc(c.expr)
       ts.NoError(err)
       ts.Equalf(c.expected, res.Value, "Expression: %s\nActual: %s\nExpected: %s", c.expr, res.Value, c.expected)
   }
}
