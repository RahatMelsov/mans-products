package structs

import (
	"math/rand"
	"time"
)

type UserInfo map[string]string
type UsersList []UserInfo
type Animals interface {
	Says() string
	NumbersOfLegs() int
	RandomIt(int) int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *Dog) NumbersOfLegs() int {
	return 4
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) RandomIt(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

func (g *Gorilla) NumbersOfLegs() int {
	return 2
}

func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) RandomIt(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
