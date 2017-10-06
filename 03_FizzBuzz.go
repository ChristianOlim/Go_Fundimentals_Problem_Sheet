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

	//This if statement solves 'FizzBuzz' for multiples 
	//of 3 and 5. It also printing 'Fizz' for muliples
	//of 3 and prints 'Buzz' for multiples of 5.
      if i%15 == 0 {
	fmt.Println("FizzBuzz")
	} else if i%3 == 0 {
	fmt.Println("Fizz")
	} else if i%5 == 0 {
	fmt.Println("Buzz")
	} else {
	fmt.Println(i)
	}
		
    }
}