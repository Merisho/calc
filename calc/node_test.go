package calc

import (
    "github.com/stretchr/testify/suite"
    "testing"
)

func TestNode(t *testing.T) {
    suite.Run(t, new(NodeTestSuite))
}

type NodeTestSuite struct {
    suite.Suite
}

func (ts *NodeTestSuite) TestSubtractNestedWithParentesesExpression() {
    r := NewNode("+")
    r.PushChild(NewNode("+", 10))

    c := NewNode("-")
    c.PushChild(NewNode("+", 5))
    c.PushChild(NewNode("+", 2))

    r.PushChild(c)

    ts.EqualValues(3, r.Calc())
}

func (ts *NodeTestSuite) TestSequentialDivision() {
    r := NewNode("+")
    r.PushChild(NewNode("+", 24))
    r.PushChild(NewNode("/", 3))
    r.PushChild(NewNode("/", 4))

    ts.EqualValues(2, r.Calc())
}
