package calc

import "github.com/merisho/calculator/calc/tokens"

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
    res := n.val
    for i := 0; i < len(n.children); i++ {
        c := n.children[i]
        if len(c.children) == 0 {
            res = c.Operator().Apply(res, c.val)
        } else {
            res += c.Calc()
        }
    }

    return int64(n.op.Sign()) * res
}

func (n *Node) Operator() tokens.Operator {
    return n.op
}

func (n *Node) signedVal() int64 {
    return int64(n.op.Sign()) * n.val
}
