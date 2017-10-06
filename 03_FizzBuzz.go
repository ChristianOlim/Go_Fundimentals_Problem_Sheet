//Source: https://tour.golang.org/welcome/4
// https://gobyexample.com/if-else
package main

//Importting fmt and time
import (
	"fmt"
)

func main() {  
	
	//Printing numbers from 1 to 100
    for i := 1; i <= 100; i++ {

		//This method solves 'FizzBuzz' for multiples of 3 and 5
        if i%15 == 0 {
            fmt.Println("FizzBuzz")
		}
		//This if statement prints 'Fizz' for multiples of 3
		else if i%3 == 0 {
            fmt.Println("Fizz")
		}
		//This if statement prints 'Buzz' for multiples of 5
		else if i%5 == 0 {
            fmt.Println("Buzz")
		}
		//This will print out the number if it's not divisible by 3 and/ or 5
		else {
            fmt.Println(i)
		}
		
    }
}