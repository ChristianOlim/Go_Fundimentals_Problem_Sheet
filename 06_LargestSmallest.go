//Source: https://adriann.github.io/programming_problems.html
//https://gist.github.com/pyk/10519339
package main

//Package fmt implements formatted I/O
import (
	"fmt"
)

func main() {
	//Here is an array of numbers to use before we
	//sort out from smallest to largest

	x := []int{
		41, 78, 57, 11,
		84, 36, 27, 68,
		15, 88, 45, 18,
		1, 22, 72, 92,
	}

	smallest, largest := x[0], x[0]
	for _, v := range x {
		if v > largest {
			largest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	fmt.Println("====== Qustion 6 ======")
	fmt.Println("The largest number is ", largest)
	fmt.Println("The smallest number is ", smallest)
}