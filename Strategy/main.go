package strategy

import "fmt"

/*
Type: Behavioral
Purpose: Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.
Additional: - allow behavior or algorithm to be changed at run time
			- avoid exposing complex, algorithm-specific data structures
			- instead of many conditionals, move related conditional branches into their own Strategy class
*/

// Context can change strategy (algorithm/behavior) at run time
type Context struct {
	strategy Strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.execute(a, b)
}

type Strategy interface {
	execute(a, b int) int
}

type MinusStrategy struct{}

func (*MinusStrategy) execute(a, b int) int {
	return a - b
}

type PlusStrategy struct{}

func (*PlusStrategy) execute(a, b int) int {
	return a + b
}

type MulStrategy struct{}

func (*MulStrategy) execute(a, b int) int {
	return a * b
}

func main() {
	ctx := &Context{}
	ctx.strategy = &PlusStrategy{} // change strategy to plus
	fmt.Println(ctx.ExecuteStrategy(3, 2))

	ctx.strategy = &MinusStrategy{} // change strategy to minus
	fmt.Println(ctx.ExecuteStrategy(3, 2))

	ctx.strategy = &MulStrategy{} // change strategy to multiply
	fmt.Println(ctx.ExecuteStrategy(3, 2))
}
