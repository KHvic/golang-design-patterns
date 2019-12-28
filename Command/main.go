package command

import "fmt"

/*
	Type: Behavioral
	Purpose: Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.
	Additional: - prevent having to hardcode many if else statements to handle commands
				- thus adding new commands does not affect existing code
				- allow implementation of redo and undo by having a stack of commands and states
*/

// Receivers are targeted entities that store states
type Calculator struct {
	value int
}

// Command perform action on some Receiver
type Command interface {
	execute()
	unexecute() // Not needed if we do not need to implement undo
}

type AddCommand struct {
	calculator *Calculator
	valToAdd   int
}

func (c *AddCommand) execute() {
	c.calculator.value += c.valToAdd
}

func (c *AddCommand) unexecute() {
	c.calculator.value -= c.valToAdd
}

type MultiplyCommand struct {
	calculator *Calculator
	valToMul   int
}

func (c *MultiplyCommand) execute() {
	c.calculator.value *= c.valToMul
}

func (c *MultiplyCommand) unexecute() {
	c.calculator.value /= c.valToMul
}

// Invoker execute commands and perform bookkeeping for undo/redo
type Invoker struct {
	doStack []Command
}

func (i *Invoker) Do(c Command) {
	i.doStack = append(i.doStack, c)
	c.execute()
}

func (i *Invoker) Redo() {
	len := len(i.doStack)
	if len > 0 {
		previousCommand := i.doStack[len-1]
		i.Do(previousCommand)
	}
}

func (i *Invoker) Undo() {
	len := len(i.doStack)
	if len > 0 {
		previousCommand := i.doStack[len-1]
		previousCommand.unexecute()
		i.doStack = i.doStack[0 : len-1]
	}
}

func main() {
	invoker := &Invoker{}
	calc := &Calculator{value: 1}
	add10 := &AddCommand{calc, 10}
	mul10 := &MultiplyCommand{calc, 10}

	invoker.Do(add10)
	fmt.Println(calc.value)
	invoker.Redo()
	fmt.Println(calc.value)
	invoker.Do(mul10)
	fmt.Println(calc.value)
	invoker.Undo()
	fmt.Println(calc.value)
}
