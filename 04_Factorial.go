//Source: https://gist.github.com/esimov/9622710
package main

//Package fmt implements formatted I/O
import (
	"fmt"
)

//Here is the method for calculating the sum of a factorial
func factorial(f uint) uint {
	if f == 0 {
		return 1
	}
	return f * factorial(f - 1)
}

//This is where I print everything to the screen
func main() {
	f := uint(10)	
	calcFactorial := factorial(f)
	fmt.Println("")
	fmt.Println("=========== Question 4 ===============")
	fmt.Println("The sum of the factorial 10 is:",calcFactorial)	
	fmt.Println("")
}