package main

import "fmt"

func main() {
	age := 18

	// Control flow using the if - else statements
	// The 'else' block is optional
	if age > 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("You're too young!")
	}

	// You can also declare a variable within an 'if' statement
	// The variable will only exist within the 'if - else' block
	if score := 85; score > 90 {
		fmt.Println("Grade A")
	} else {
		fmt.Println("Grade B")
	}

	// In Go, looping is only allowed using the 'for' statment
	// even if we want to loop in a 'while' style like in other
	// programming languages
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// A small demo using 'for' and 'if' in combination
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz:", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz:", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz:", i)
		} else {
			fmt.Println(i)
		}
	}
}
