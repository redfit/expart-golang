package main

import "fmt"

type CryFootstepper interface {
	Crier
	Footstepper
}

type Crier interface {
	Cry() string
}

type Footstepper interface {
	Footsteps() string
}

type Person struct{}

func (p *Person) Cry() string {
	return "Hi"
}

func (p *Person) Footsteps() string {
	return "Pitapat"
}

type PartyPeople struct {
	Person
}

func (p *PartyPeople) Cry() string {
	return "Sup?"
}

func main() {
	var cf CryFootstepper

	cf = &Person{}
	fmt.Println(cf.Cry(), cf.Footsteps())

	cf = &PartyPeople{}
	fmt.Println(cf.Cry(), cf.Footsteps())
}
