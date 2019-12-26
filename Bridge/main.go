package bridge

import "fmt"

/*
	Type: Structural
	Purpose: Decouple an abstraction from its implementation so that the two can vary independently.
			 Allow splitting a large class into two indepedent & separated hierarchies
	Additional: - for avoiding a permanent binding between an abstraction and its implementation. For example, when the implementation must be selected or switched at run-time.
				- allow sharing an implementation among multiple objects
*/

// Abstraction hierarchy with its operation
type Weapon interface {
	swing() // operation
	getEnchantment() Enchantment
}

// Refined abstractions
type Sword struct {
	enchantment Enchantment
}

func (s *Sword) getEnchantment() Enchantment {
	return s.enchantment
}
func (s *Sword) swing() {
	fmt.Println("The sword is swinged.")
	s.enchantment.apply()
}

type Hammer struct {
	enchantment Enchantment
}

func (h *Hammer) getEnchantment() Enchantment {
	return h.enchantment
}
func (h *Hammer) swing() {
	fmt.Println("The hammer is swinged.")
	h.enchantment.apply()
}

// Implementor hierarchy
type Enchantment interface {
	apply()
}

// Concrete Implementors
type FireEnchant struct{}

func (f *FireEnchant) apply() {
	fmt.Println("The item shoot fireballs at the enemies")
}

type ShockEnchant struct{}

func (s *ShockEnchant) apply() {
	fmt.Println("The item summon lightnings at the enemies")
}

func main() {
	fireSword := &Sword{&FireEnchant{}}
	shockHammer := &Hammer{&ShockEnchant{}}
	fireSword.swing()
	shockHammer.swing()
}
