package main

import "fmt"

type Entety interface {
	toString() string
}

//Person
type Person struct {
	name  string
	age   int
	cpf   string
	email string
}

func (p *Person) setName(value string) {
	p.name = value
}

func (p *Person) setAge(value int) {
	p.age = value
}

func newPerson(name string) *Person {
	p := Person{name: name, age: 0}
	return &p
}

//Car
type Car struct {
	name string
	user Person
}

func (c *Car) setName(value string) {
	c.name = value
}

func (c *Car) setUser(newUser Person) {
	c.user = newUser
}

func main() {
	person1 := Person{name: "Charles", age: 21, cpf: "017.626.944-41"}
	car1 := Car{name: "Duster"}

	car1.setUser(person1)
	fmt.Println(car1.user.name)

	person1.setName("Charles Lindo")
	fmt.Println(person1.name)
	fmt.Println(car1.user.name)
}
