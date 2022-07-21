package main

func main() {
	/* var p = Person{"John", "Doe"} // first way*/

	/* p := Person{firstName: "John", lastName: "Doe"} //second way */

	var p person // third way
	/* fmt.Printf("%+v\n", p) */
	p.firstName = "John"
	p.lastName = "Doe"
	p.contactInfo.email = "email.com"
	p.contactInfo.zipCode = 12345
	//p.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	/* (&jim).updateName("jimmy") */
	jim.print()
}
