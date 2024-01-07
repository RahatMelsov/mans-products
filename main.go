package main

import (
	"fmt"

	structs "github.com/RahatMelsov/mansProducts/contstants"
)

func main() {
	dog := structs.Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	gorilla := structs.Gorilla{
		Name:          "Samson",
		Color:         "Black",
		NumberOfTeeth: 4,
	}
	PrintInfo(&gorilla)
	PrintInfo(&dog)
}

func PrintInfo(a structs.Animals) {
	fmt.Println("this animals says", a.Says(), "and has", a.NumbersOfLegs(), "legs")
}
