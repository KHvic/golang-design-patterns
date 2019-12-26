package adapter

import "fmt"

/*
	Type: Structural
	Purpose: Convert the interface of a class into another interface the clients expect.
			 Also known as wrapper, allowing incompatible classes to work together.
	Additional: - Global variables are bad because it can be affected by code from anywhere
				- Global shared instance which prevents an object and resources used by this object from being deallocated
				- Creates tightly coupled code. The clients of the Singleton become difficult to test
*/

// Adaptee concrete
type Adaptee struct{}

func (a Adaptee) SpecificRequest() {
	fmt.Println("Adaptee's Request")
}

// Target Interface
type Target interface {
	Request()
}

// Adapter concrete class implementing Target
// Basically a wrapper to make Adaptee conform to Target class requirement of having Request method
type Adapter struct {
	adaptee Adaptee
}

func (a Adapter) Request() {
	a.adaptee.SpecificRequest()
}

func main() {
	adaptee := Adaptee{}
	var target Target = Adapter{adaptee}
	// can now use Adaptee methods via Target
	target.Request()
}
