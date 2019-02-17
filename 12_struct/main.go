package main

import (
	"fmt"
	"strconv"
)

// Person Struct defination
type Person struct {
	// firstName string
	// lastName  string
	// email     string
	// age       int

	// you can also declare this way
	firstName, lastName, email string
	age                        int
}

// Struct function (value receiver)
func (p Person) greetings() string {
	return "Greetings, " + p.firstName + " " + p.lastName
}

func (p Person) findAge() string {
	return "Your age is " + strconv.Itoa(p.age)
}

// Struct function (Pointer receiver)
func (p *Person) changeName(firstName, lastName string) {
	p.firstName = firstName
	p.lastName = lastName
}

func (p *Person) changeAge() {
	p.age++
}

func main() {
	// init person struct
	person1 := Person{firstName: "Mike", lastName: "Tyson", email: "mike@gmail.com", age: 45}
	fmt.Println(person1)

	// alternative way to initi person struct
	person2 := Person{"Terry", "benedict", "terry@gmail.com", 28}
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person1.age = 16
	fmt.Println(person1)

	// printing firstname using struct function
	fmt.Println(person2.greetings())
	fmt.Println(person1.findAge())

	// changing values using Struct pointer function
	person1.changeName("Micheal", "Jackson")
	person1.changeAge()
	fmt.Println(person1)
}
