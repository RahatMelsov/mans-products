package structs

type UserInfo map[string]string
type UsersList []UserInfo
type Animals interface {
	Says() string
	NumbersOfLegs() int
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

func (g *Gorilla) NumbersOfLegs() int {
	return 2
}

func (g *Gorilla) Says() string {
	return "Ugh"
}
