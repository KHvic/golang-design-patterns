package builder

import "fmt"

/*
	Type: Creational
	Purpose: Separate the construction of a complex object from its representation so that the same construction process can create different representations.
	Additional: - allow construct complex objects step by step
				- help avoid having a object constructor with many parameters
*/
// Example of creating a game hero

type Profession string

const (
	Mage    Profession = "Mage"
	Warrior Profession = "Warrior"
	Archer  Profession = "Archer"
)

type Weapon string

const (
	Knife Weapon = "Knife"
	Staff Weapon = "Staff"
	Bow   Weapon = "Bow"
)

type Armor string

const (
	Leather Armor = "Leather"
	Steel   Armor = "Steel"
)

type Color string

const (
	Red    Color = "Red"
	Black  Color = "Black"
	Yellow Color = "Yellow"
)

type Hero struct {
	profession  Profession
	weapon      Weapon
	weaponColor Color
	armor       Armor
	armorColor  Color
}

type Builder struct {
	hero Hero
}

func (b *Builder) withWeapon(w Weapon, c Color) *Builder {
	b.hero.weapon = w
	b.hero.weaponColor = c
	return b
}

func (b *Builder) withArmor(a Armor, c Color) *Builder {
	b.hero.armor = a
	b.hero.armorColor = c
	return b
}

func NewBuilder(p Profession) *Builder {
	return &Builder{Hero{profession: p}}
}

func (b *Builder) build() *Hero {
	return &b.hero
}

func main() {
	mageHero := NewBuilder(Mage).withArmor(Leather, Red).withWeapon(Staff, Black).build()
	fmt.Println(*mageHero)
}
