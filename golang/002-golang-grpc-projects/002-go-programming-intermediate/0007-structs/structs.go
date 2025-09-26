package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	age int
	// Embedded struct
	address Address
	// Embedded Anonymous struct
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
}

// Embed struct
type Address struct {
	city string
	country string
}

// Method - value receiver
func (p Person) fullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

// Method - pointer receiver
func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {
	p1 := Person{
		firstName: "John",
		lastName: "Don",
		age: 34,
		address: Address{
			city: "Bangalore",
			country: "India",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "2349999999",
			cell: "888898765",
		},
	}
	fmt.Println("First Name of person 1:", p1.firstName)
	fmt.Println("Last Name of person 1:", p1.lastName)
	fmt.Println("Age of person 1:", p1.age)
	fmt.Println("Full Name of person 1", p1.fullName())
	fmt.Println("City of person 1:", p1.address.city)
	fmt.Println("Country of person 1:", p1.address.country)
	// Anonymous Fields
	fmt.Println("Home phone number of person 1:", p1.home)
	fmt.Println("Cell number of person 1:", p1.cell)

	p2 := Person{
		firstName: "Dev",
	}
	// Another way to add a value to embedded struct
	p2.address.city = "New York"
	p2.address.country = "USA"

	p3 := Person{
		firstName: "Dev",
		address: Address{
			city: "New York",
			country: "USA",
		},
	}

	// Compare structs
	fmt.Println("Are p1 and p2 equal:", p1 == p2)
	fmt.Println("Are p2 and p3 equal:", p2 == p3)

	fmt.Println("First Name of person 2:", p2.firstName)
	fmt.Println("Last Name of person 2:", p2.lastName)
	fmt.Println("Before increment - age of person 2:", p2.age)
	p2.incrementAgeByOne()
	fmt.Println("After increment - age of person 2:", p2.age)
	fmt.Println("Full Name of person 2", p2.fullName())
	fmt.Println("City and Country:", p2.address)
	fmt.Println("City of person 2:", p2.address.city)
	fmt.Println("Country of person2:", p2.address.country)

	// Anonymous structs
	user := struct {
		userName string
		userEmail string
	} {
		userName: "User1234",
		userEmail: "user1234@gmail.com",
	}
	fmt.Println("User Name:", user.userName)
	fmt.Println("Email ID:", user.userEmail)
}

/*
------------------Output-----------------
First Name of person 1: John
Last Name of person 1: Don
Age of person 1: 34
Full Name of person 1 John Don
City of person 1: Bangalore
Country of person 1: India
Home phone number of person 1: 2349999999
Cell number of person 1: 888898765
Are p1 and p2 equal: false
Are p2 and p3 equal: true
First Name of person 2: Dev
Last Name of person 2: 
Before increment - age of person 2: 0
After increment - age of person 2: 1
Full Name of person 2 Dev 
City and Country: {New York USA}
City of person 2: New York
Country of person2: USA
User Name: User1234
Email ID: user1234@gmail.com
------------------Output-----------------
*/