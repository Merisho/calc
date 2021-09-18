package calc

import "github.com/merisho/calculator/calc/tokens"

type NodeStack struct {
    nodes []*Node
}

func (s *NodeStack) Push(n *Node) {
    s.nodes = append(s.nodes, n)
}

func (s *NodeStack) Top() *Node {
    return s.nodes[len(s.nodes) - 1]
}

func (s *NodeStack) Pop() {
    s.nodes = s.nodes[:len(s.nodes) - 1]
}

type OperatorStack struct {
    ops []tokens.Operator
}

func (s *OperatorStack) Push(o tokens.Operator) {
    s.ops = append(s.ops, o)
}

func (s *OperatorStack) Top() tokens.Operator {
    return s.ops[len(s.ops) - 1]
}

func (s *OperatorStack) Pop() {
    s.ops = s.ops[:len(s.ops) - 1]
}
