package strategy

// IsStrategy
// @Description: 定义一个策略类
type IsStrategy interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IsStrategy
}

// setStrategy
// @Description: 设置策略
// @receiver operator
// @param strategy
func (operator *Operator) setStrategy(strategy IsStrategy) {
	operator.strategy = strategy
}

// calculate
// @Description: 调用策略中的方法
// @receiver operator
// @param a
// @param b
// @return int
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
