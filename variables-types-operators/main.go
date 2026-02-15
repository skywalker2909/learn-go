package main

import "fmt"

func main() {
	// This is the short way of declaring variables using the walrus
	// operator ( := ) wherego determines the type based on the value
	// (type inference)
	// Note: You cannot use the walrus operator outside of a function.
	firstName := "Neville"
	age := 29
	isLearning := true

	fmt.Println("Name:", firstName)
	fmt.Println("Age:", age)
	fmt.Println("IsLearning:", isLearning)

	// This is the formal way of declaring variables where the type of
	// value is explicitly specified.
	// You can also use this when you want to define variables without
	// aninitial value or you want to define them globally outside of
	// the function.
	var salary float64 = 100000.9999
	fmt.Println("Salary:", salary)

	var unknownNumber int
	var isKnown bool
	var defaultValue string

	// Variables have a 'zero value' by default in Go compared to other
	// languages where it's a null or undefined.
	fmt.Println("UnknownNumber:", unknownNumber)
	fmt.Println("IsKnown:", isKnown)
	fmt.Println("Default Value:", defaultValue)

	price := 100.99
	quantity := 3

	// Go does not allow to multiply an int and float directly. You have
	// to convert the int to a float.
	total := price * float64(quantity)
	fmt.Println("The total cost is:", total)
}
