package composite

import "fmt"

/*
   Type: Structural
   Purpose: Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.
   Additional: - Represent hierarchies of objects
            - Ignore the difference between compositions of objects and individual objects
*/

type Component interface {
	printInfo() // operation()
}

// Composite containing more components
type Department struct {
	members []Component
}

func (d *Department) add(c Component) {
	d.members = append(d.members, c)
}

func (d *Department) remove(c Component) {
	var res []Component
	for _, com := range d.members {
		if com != c {
			res = append(res, com)
		}
	}
	d.members = res
}

func (d *Department) printInfo() {
	for _, com := range d.members {
		com.printInfo()
	}
}

// Leaf component
type Employee struct {
	name string
}

func (e *Employee) printInfo() {
	fmt.Println("Name:", e.name)
}

func main() {
	engineeringDept := &Department{}
	engineeringDept.add(&Employee{"John"})
	engineeringDept.add(&Employee{"Mary"})
	salesDept := &Department{}
	salesDept.add(&Employee{"Ken"})
	salesDept.add(&Employee{"Lizzy"})
	company := &Department{[]Component{engineeringDept, salesDept}}

	company.printInfo()
}
