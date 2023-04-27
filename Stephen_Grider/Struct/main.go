package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	// alex := person{"alex", "anderson"} /*option 1 on declaring struct*/
	// alex := person{firstName: "alex", lastName: "anderson"} /*option 2 on declaring struct*/
	// var alex person /*option 3 on declaring struct*/
	// alex.firstName = "Alex" /*option 3 on declaring struct*/
	// alex.lastName = "Anderson"/*option 3 on declaring struct*/
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@gmail.com",
		// 	zipcode: 94000,
		// },
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipcode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
	// fmt.Printf("%+v", jim)

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
