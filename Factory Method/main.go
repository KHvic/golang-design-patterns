package factory_method

import "fmt"

/*
	Type: Creational
	Purpose: Provides an interface for creating an instance of a class, with its subclasses deciding which class to instantiate.
	Additional: - use when we cannot anticipate which class to instantiate (for example, allow client's input to decide which class to create)
				- to delegate responsbilities of creating class
*/

type CardGame interface {
	StartGame()
}

type PokerGame struct{}

func (p PokerGame) StartGame() {
	fmt.Println("Starting a Poker Game")
}

type BlackjackGame struct{}

func (b BlackjackGame) StartGame() {
	fmt.Println("Starting a Blackjack Game")
}

type UnoGame struct{}

func (u UnoGame) StartGame() {
	fmt.Println("Starting a UnoGame Game")
}

// CreateCardGame is a factory method that takes in a type and return an instance of it
func CreateCardGame(t string) CardGame {
	switch t {
	case "POKER":
		return PokerGame{}
	case "BLACKJACK":
		return BlackjackGame{}
	default:
		return UnoGame{}
	}
}

func main() {
	blackjack := CreateCardGame("BLACKJACK")
	blackjack.StartGame()
}
