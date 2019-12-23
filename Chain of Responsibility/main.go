package chain

import (
	"fmt"
)

/*
Type: Behavioral
Purpose: Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request.
Chain the receiving objects and pass the request along the chain until an object handles it.
Additional: - when more than one object may handle a request, and the handler isn't known beforehand
   - you want to issue a request to one of several objects without specifying the receiver explicitly
*/

const (
	CommanderOnly = iota
	OfficerOnly
	SoldierOnly
)

type Request struct {
	Message     string
	RequestType int
}

type RequestHandler interface {
	HandleRequest(*Request)
}

type Commander struct {
	Officer *Officer
}

type Officer struct {
	Soldier *Soldier
}

type Soldier struct{}

func (c *Commander) HandleRequest(r *Request) {
	if CommanderOnly == r.RequestType {
		fmt.Println("Commander is handling request: ", r.Message)
	} else {
		c.Officer.HandleRequest(r)
	}
}

func (o *Officer) HandleRequest(r *Request) {
	if OfficerOnly == r.RequestType {
		fmt.Println("Officer is handling request: ", r.Message)
	} else {
		o.Soldier.HandleRequest(r)
	}
}

func (s *Soldier) HandleRequest(r *Request) {
	if SoldierOnly == r.RequestType {
		fmt.Println("Soldier is handling request: ", r.Message)
	}
}

func main() {
	var chain RequestHandler = &Commander{Officer: &Officer{Soldier: &Soldier{}}}
	chain.HandleRequest(&Request{Message: "Charge!!", RequestType: SoldierOnly})
	chain.HandleRequest(&Request{Message: "Train soldiers", RequestType: OfficerOnly})
	chain.HandleRequest(&Request{Message: "Plan for war", RequestType: CommanderOnly})
}
