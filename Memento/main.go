package memento

import "fmt"

/*
   Type: Behavioral
   Purpose: Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.
   Additional: - a snapshot of an object's state must be saved so that it can be restored to that state later
*/

// GameState
type State struct {
	score      int
	timePlayed int
}

// Originator with actual state
type GameOriginator struct {
	state State
}

func (o *GameOriginator) createMemento() *GameMemento {
	m := &GameMemento{}
	m.setState(o.state)
	return m
}

func (o *GameOriginator) setMemento(m *GameMemento) {
	o.state = m.getState()
}

// Memento with a saved copy of the state
type GameMemento struct {
	state State
}

func (m *GameMemento) setState(s State) {
	m.state = s
}
func (m *GameMemento) getState() State {
	return m.state
}

func main() {
	originator := &GameOriginator{State{100, 10}}
	memento := originator.createMemento() // save game
	fmt.Println(originator.state)

	// loss large score in the game, so we want to rollback
	originator.state.score -= 100
	originator.state.timePlayed += 20
	fmt.Println(originator.state)

	originator.setMemento(memento)
	fmt.Println(originator.state)
}
