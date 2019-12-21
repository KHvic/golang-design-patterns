package facade

import "fmt"

/*
	Type: Structural
	Purpose: Provides a simplified & unified interface to complex subsystems. Facade defines a higher-level interface that makes the subsystems easier to use.
	Additional: Hide complexities of underlying subsystems
*/

// Consider the case that the Client want to run many multiple subsystems together without handling the complexity of these systems

// Subsystem that we want to hide away from our client
type Subsystem interface {
	Run()
}

type SubsystemA struct{}

func (*SubsystemA) Run() {
	fmt.Println("SubsystemA is running some complex operation!")
}

type SubsystemB struct{}

func (*SubsystemB) Run() {
	fmt.Println("SubsytemB is running some complex operation!")
}

// Create a Facade to expose unified & simplied API to client
type Facade struct {
	subsystems []Subsystem
}

func NewFacade() *Facade {
	fmt.Println("Setting up certain configuration settings for subsystems")
	subsystems := []Subsystem{&SubsystemA{}, &SubsystemB{}}
	return &Facade{subsystems}
}

// API to simply run all underlying subsytems
func (f *Facade) ExecuteAll() {
	for _, s := range f.subsystems {
		s.Run()
	}
}
