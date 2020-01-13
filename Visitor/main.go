package visitor

import "fmt"

/*
   Type: Behavioral
   Purpose: Represent an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operates.
   Additional: - Move operation logic to another class (Visitor class)
			   - Allow defining of new operations over the structure, don't need to modify Element
			   - Perform operations depending on concrete classes
*/

// Elements are object that can hold states
type Element interface {
	accept(visitor Visitor)
}

type ElementA struct{}

func (e *ElementA) accept(visitor Visitor) {
	visitor.visitA(e)
}

type ElementB struct{}

func (e *ElementB) accept(visitor Visitor) {
	visitor.visitB(e)
}

// Visitors (Classes holding the actual operations for each element)
type Visitor interface {
	visitA(elem *ElementA)
	visitB(elem *ElementB)
}

type OldVisitor struct{}

func (*OldVisitor) visitA(*ElementA) {
	fmt.Println("Running Old Operation for Element A")
}
func (*OldVisitor) visitB(*ElementB) {
	fmt.Println("Running Old Operation for Element B")
}

type NewVisitor struct{}

func (*NewVisitor) visitA(*ElementA) {
	fmt.Println("Running New Operation for Element A")
}
func (*NewVisitor) visitB(*ElementB) {
	fmt.Println("Running New Operation for Element B")
}

func main() {
	elemA := &ElementA{}
	elemB := &ElementB{}
	oldOperations := &OldVisitor{}
	newOperations := &NewVisitor{}
	oldOperations.visitA(elemA)
	oldOperations.visitB(elemB)
	newOperations.visitA(elemA)
	newOperations.visitB(elemB)
}
