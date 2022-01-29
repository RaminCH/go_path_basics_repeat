package main

import "fmt"

//-----------Interface for animals
type Animal interface {
	Says() string
	NunmberOfLegs() int
}

//------------Initializing each animal struct
type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}



func main() {
	dog := Dog{
		Name: "Samsun",
		Breed: "German Shepherd",
	}

	gorilla := Gorilla{
		Name: "Jock",
		Color: "Grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&dog)
	PrintInfo(&gorilla)
}

//------------------------------Creating common print info func
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NunmberOfLegs(), "legs")
}


//---------------------------------Giving methods from Animal interface to each animal
func(d *Dog) Says() string {
	return "woof"
}

func(d *Dog) NunmberOfLegs() int {
	return 4
}

func(g *Gorilla) Says() string {
	return "ugh"
}

func(g *Gorilla) NunmberOfLegs() int {
	return 2
}

//Output

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Interfaces> go run main.go
// This animal says woof and has 4 legs
// This animal says ugh and has 2 legs