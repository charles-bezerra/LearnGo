package main

import "fmt"

type Person struct {
	name string
	age int
	cpf string
	email string
}

func (p *Person) setName(value string) {
	p.name = value
}	

func (p *Person) setAge(value int) {
	p.age = value
}

func newPerson(name string) *Person {
	p := Person{name: name, age: 12}
	return &p
}


type Car struct {
	name string
}

func (c *Car) setName(value string) {
	c.name = value
}

func main() {
	p := newPerson("Charles")
	p.setName("Charles Bezerra")
	p.setAge(21)
	fmt.Println(p.age)

	c := Car{name: "car"}
	c.setName("Car")
	fmt.Println(c.name)
}
