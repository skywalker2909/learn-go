// This tells the go compiler that this file should compile into an executable program
// instead of a shared library.
package main

// Go has a massive standard library. "fmt" stands for Format. It handles input/output.
import "fmt"

// This is the entry point. When you run a program. When you run a program, Go looks for
// 'main' and if it's not present then this program wont run.
func main() {
	fmt.Println("Hello beautiful world!")
	fmt.Println("Welcome to the fun and adventurous world of golang!")
}
