package main

import (
	"fmt"
)

// var email string = "reungthanoo_r@silpakorn.edu"

func main() {
	// var name string = "ruthaichanok"
	var age int = 23

	email := "reungthanoo_r@silpakorn.edu"
	gpa := 2.62

	firstname, lastname := "ruthaichanok", "reungthanoo"
	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)
}
