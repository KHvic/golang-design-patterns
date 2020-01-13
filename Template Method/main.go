package template_method

import "fmt"

/*
   Type: Behavioral
   Purpose: Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.
   Additional: - When two components have similarities but have part of the implementation that differs. Use Template Method to make part of the implementation reusable or different.
               - Allow variation of a algorithm/behavior
*/

// Since Golang doesn't support inheritance, we can just define an interface composition on an object
// This make it like a "Strategy Pattern"
// In original implementation, the template classes are concrete class that inherit from partially defined abstract class
type Car struct {
	fuel     int
	distance int
	wheels   Wheels
}

// Defer part of the operations to be defined by the template classes
func (c *Car) Drive() {
	if c.fuel > 0 {
		c.fuel--
		c.distance += c.wheels.move()
		fmt.Printf("Car is at distance: %v\n", c.distance)
	} else {
		fmt.Printf("Car is out of fuel!\n")
	}
}

// Template
type Wheels interface {
	move() int
}

// Concrete implementations defining the operation
type BrokenWheels struct{}

func (*BrokenWheels) move() int {
	return 0
}

type CheapWheels struct{}

func (*CheapWheels) move() int {
	return 10
}

type ExpensiveWheels struct{}

func (*ExpensiveWheels) move() int {
	return 30
}

func main() {
	car := &Car{3, 0, &CheapWheels{}}
	car.Drive()
	car.wheels = &BrokenWheels{}
	car.Drive()
	car.wheels = &ExpensiveWheels{}
	car.Drive()
	car.Drive()
}
