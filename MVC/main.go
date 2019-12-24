package mvc

import "fmt"

/*
Type: Architectural
Purpose: Separate the user interface into three interconnected components: the model, the view and the controller.
Let the model manage the data, the view display the data and the controller mediate updating the data and redrawing the display.
Additional: - Allow Model to be reused, can have multiple different view (View can also made reusable by abstracting Model)
   - Clear separation of logic; separate the domain data from its user interface representation
*/

// Model only represent data
type Model struct {
	name     string
	color    string
	quantity int
}

// View
type View struct{}

func (*View) Display(m *Model) {
	if m.quantity == 0 {
		fmt.Println("Nothing is displayed")
	} else {
		fmt.Printf("Displaying %d %s colored %s\n", m.quantity, m.color, m.name)
	}
}

// Controller separates depedency between view and model
type Controller struct {
	model *Model
	view  *View
}

func (c *Controller) add() {
	c.model.quantity++
}

func (c *Controller) dec() {
	c.model.quantity--
	if c.model.quantity < 0 {
		c.model.quantity = 0
	}
}

func (c *Controller) setColor(color string) {
	c.model.color = color
}

func (c *Controller) setName(name string) {
	c.model.name = name
}

func (c *Controller) updateView() {
	c.view.Display(c.model)
}

func main() {
	model := &Model{name: "Chameleon", color: "Gray", quantity: 2}
	view := &View{}
	controller := &Controller{model, view}
	controller.updateView()
	controller.dec()
	controller.setColor("Purple")
	controller.updateView()
}
