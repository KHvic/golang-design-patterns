package observer

import "fmt"

/*
	Type: Behavioral
	Purpose: Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
	Additional: - When an object should be able to notify other objects without making assumptions about who these objects are.
				- Use when we do not want observers to continuously poll for an object state change.
*/

type Subject struct {
	state     int
	observers []Observer
}

// make observer subject to the subject
func (s *Subject) registerObserver(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) unregisterObserver(o Observer) {
	newObservers := make([]Observer, 0, len(s.observers))
	for _, cur := range s.observers {
		if cur != o {
			newObservers = append(newObservers, o)
		}
	}
	s.observers = newObservers
}

// notify all observers subscribed to the subject
func (s *Subject) notifyObservers() {
	for _, o := range s.observers {
		o.update(s.state)
	}
}

type Observer interface {
	update(int)
}

type ConcreteObserverA struct{}

func (o *ConcreteObserverA) update(state int) {
	fmt.Println("ConcreteObserverA updated ", state)
}

type ConcreteObserverB struct{}

func (o *ConcreteObserverB) update(state int) {
	fmt.Println("ConcreteObserverB updated ", state)
}

func main() {
	subject := &Subject{state: 1}
	o1 := &ConcreteObserverA{}
	o2 := &ConcreteObserverB{}

	subject.registerObserver(o1)
	subject.registerObserver(o2)
	subject.notifyObservers() // should notify o1 and o2 with state = 1

	subject.unregisterObserver(o1)
	subject.state = 2
	subject.notifyObservers() // should notify o2 with state = 2
}
