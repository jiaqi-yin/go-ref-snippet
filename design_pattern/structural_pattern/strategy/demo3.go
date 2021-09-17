package strategy

type IStrategy interface {
	do(int, int) int
}

type addition struct{}

func (*addition) do(a, b int) int {
	return a + b
}

type subtraction struct{}

func (*subtraction) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
