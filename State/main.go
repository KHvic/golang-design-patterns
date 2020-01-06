package state

/*
	Type: Behavioral
	Purpose: Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.
	Additional: - An object's behavior depends on its state, and it must change its behavior at run-time depending on that state.
				- Instead of having many conditional statement, consider using this design, putting each condition on each state.
*/

import "fmt"

type State interface {
	handle(c *Context)
}

// Alternating between the two state
type StartState struct{}

func (s *StartState) handle(c *Context) {
	fmt.Println("Moving to StopState")
	c.setState(&StopState{})
}

type StopState struct{}

func (s *StopState) handle(c *Context) {
	fmt.Println("Moving to StartState")
	c.setState(&StartState{})
}

type Context struct {
	state State
}

func (c *Context) setState(s State) {
	c.state = s
}

func (c *Context) handle() {
	c.state.handle(c)
}

func main() {
	ctx := &Context{&StartState{}}
	ctx.handle()
	ctx.handle()
	ctx.handle()
}
