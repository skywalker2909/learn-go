package main

import "fmt"

// This is a function in Go that takes two integers and returns
// one integer.
// Note: if the parameters are of the same type, then you can just
// shorten it to '(a, b int)'
func add(a int, b int) int {
	return a + b
}

// In Go, functions can also return multiple values.
func divide(a, b int) (int, int) {
	result := a / b
	remainder := a % b

	return result, remainder
}

func greet(name string) string {
	return "Welcome, " + name
}

func checkAge(age int) (int, bool) {
	futureAge := age + 5
	isAdult := age >= 18

	return futureAge, isAdult
}

func main() {
	// A basic function
	message := greet("Nev")
	fmt.Println(message)

	// Another example of a basic function
	total := add(6, 7)
	fmt.Println("Total:", total)

	// A function that returns multiple values
	myCurrentAge := 13
	ageInFuture, amIAnAdult := checkAge(myCurrentAge)

	fmt.Println("In 5 years, my age will be:", ageInFuture)
	fmt.Println("Am i an adult now?:", amIAnAdult)

	// Sometimes, we may only care about one return value from the function
	// and so we can ignore all the remaining returned values using the
	// underscore symbol to throw them away because Go wont let you declare
	// the variable and not use it.
	quotient, _ := divide(11, 2)
	fmt.Println("The quotient is:", quotient)
}
