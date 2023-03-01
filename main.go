package main

import "fmt"

// Animal defines the interface for type Animal
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Dog defines the dog type
type Dog struct {
	Name  string
	Breed string
}

// Gorilla defines the Gorilla type
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

// PrintInfo recibe un valor que es de tipo Animal (la interfaz)
// Y esta interfaz tiene dos métodos: Says() y NumberOfLegs(),
// Por lo que el valor que recibe PrintInfo debe tener esos dos métodos

// Para ello se debe agregar a la estructura Dog y Gorilla
// los métodos Says() y NumberOfLegs()

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// Agregar los métodos Says() y NumberOfLegs() a la estructura Dog y Gorilla

// Podria omitirse el uso de punteros en los métodos de la interfaz,
// pero es buena práctica usarlos para evitar copias de la estructura etc.

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
