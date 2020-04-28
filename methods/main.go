package main

import "fmt"

type Person struct {
	name string
	age int
	cpf string
	email string
}

func (p *Person) setNamePerson(value string) {
	p.name = value
}	

func newPerson(name string) *Person {
	p := Person{name: name}
	return &p
}

func main() {
	p := newPerson("Charles")
	p.setNamePerson("Charles Bezerra")
	
	fmt.Printf(p.name)
}
