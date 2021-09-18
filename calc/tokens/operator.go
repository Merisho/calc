package tokens

var priority = map[Operator]int{
    "*": 2,
    "/": 2,
    "+": 1,
    "-": 1,
}

type Operator string

func (o Operator) Apply(a, b int64) int64 {
    switch o {
    case "+":
        return a + b
    case "-":
        return a - b
    case "*":
        return a * b
    case "/":
        return a / b
    }

    return 0
}

func (o Operator) Priority() int {
    return priority[o]
}

func (o Operator) Sign() int {
    if o == "-" {
        return -1
    }

    return 1
}
