package calc

import (
    "github.com/merisho/calculator/calc/tokens"
    "strconv"
    "strings"
)

func NewCalc() Calc {
    return Calc{}
}

type Calc struct {}

func (c Calc) Calc(expr string) (Result, error) {
    expr = c.prepareExpression(expr)
    toknz := tokens.Tokenize(expr)

    var tkns []tokens.Token
    for toknz.Next() {
        if toknz.Err() != nil {
            break
        }

        tkns = append(tkns, toknz.Curr())
    }

    if toknz.Err() != nil {
        return Result{}, toknz.Err()
    }

    rootNode, _ := c.buildTree(tkns, 0, "+")

    return Result{
        Value: strconv.FormatInt(rootNode.Calc(), 10),
    }, nil
}

func (c Calc) buildTree(tkns []tokens.Token, start int, op tokens.Operator) (root *Node, end int) {
    root = NewNode(op)
    node := root
    lastOp := tokens.Operator("+")
    for i := start; i < len(tkns); i++ {
        t := tkns[i]

        if n, ok := t.Number(); ok {
            node.PushChild(NewNode(lastOp, n))
        } else if op, ok := t.Operator(); ok {
            if op.Priority() > lastOp.Priority() {
                child := NewNode("+")

                child.PushChild(node.PopChild())
                node.PushChild(child)

                node = child
            } else if op.Priority() < lastOp.Priority() {
                node = root
            }

            lastOp = op
        } else if par, ok := t.Parenthesis(); ok {
            if par.Open() {
                child, end := c.buildTree(tkns, i + 1, lastOp)
                node.PushChild(child)
                i = end
            } else {
                return root, i
            }
        }
    }

    return root, len(tkns)
}

func (c Calc) prepareExpression(e string) string {
    return strings.ReplaceAll(e, " ", "")
}

type Result struct {
    Value string
}
