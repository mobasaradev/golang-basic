package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Contact struct {
	email string
	phone string
}

type Address struct {
	road     int
	house_no int
	city     string
}

type Employee struct {
	employee_data    Person
	employee_contact Contact
	employee_address Address
}

func main() {
	// First methods
	var person Person
	person.name = "Sara"
	person.age = 22
	fmt.Println("Person data is :", person)

	// Second methods
	user := Person{
		name: "Henna",
		age:  18,
	}
	fmt.Printf("Name : %v, Age : %v \n", user.name, user.age)

	// Third methods
	newPerson := new(Person)
	newPerson.name = "Evan"
	newPerson.age = 22
	fmt.Printf("Name : %v, Age : %v \n", newPerson.name, newPerson.age)

	// Nested struct
	employee := Employee{
		employee_data: Person{
			name: "Sara",
			age:  22,
		},
		employee_contact: Contact{
			email: "contact@gmail.com",
			phone: "129332434",
		},
		employee_address: Address{
			road:     03,
			house_no: 7,
			city:     "Dhaka",
		},
	}

	fmt.Printf("Employee Details : %v \n", employee)
}
