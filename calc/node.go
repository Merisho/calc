package calc

import (
    "github.com/merisho/calculator/calc/tokens"
)

func NewNode(op tokens.Operator, val ...int64) *Node {
    v := int64(0)
    if len(val) > 0 {
        v = val[0]
    }

    n := &Node{
        op:  op,
        val: v,
    }
    return n
}

type Node struct {
    children []*Node
    val int64
    op tokens.Operator
}

func (n *Node) PushChild(c *Node) *Node {
    n.children = append(n.children, c)
    return n
}

func (n *Node) PopChild() *Node {
    l := len(n.children) - 1
    c := n.children[l]

    n.children = n.children[:l]

    return c
}

func (n *Node) Calc() int64 {
    return int64(n.op.Sign()) * n.calc()
}

func (n *Node) calc() int64 {
    if len(n.children) == 0 {
        return n.val
    }

    var res int64
    for _, c := range n.children {
        res = c.Operator().Apply(res, c.calc())
    }

    return res
}

func (n *Node) Operator() tokens.Operator {
    return n.op
}
