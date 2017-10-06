//Source: https://adriann.github.io/programming_problems.html
// http://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html
package main

//Package fmt implements formatted I/O
import (
    "fmt"
    "math/rand"
    "time"
)

//This will generate a random number between a range of numbers
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := xrand(1, 10)
    guessesTaken := 0
    var guess int

    fmt.Println("================== Question 5 =======================")
    
    //this is the while loop
    for guessesTaken < 5 {
        fmt.Println("Guess a number between 1 and 10")
        fmt.Scanf("%d", &guess)
        guessesTaken++
        if guess < myrand {
            fmt.Println("Your guess is too low. Sorry")
        }
        if guess > myrand {
            fmt.Println("Your guess is too high.")
        }
        if guess == myrand {
            break
        }
    }
    //Below will print out a message to the viewer
    if guess == myrand {
        fmt.Printf("Congratulations! You guessed the number in %d tries\n", guessesTaken)
    } else {
        fmt.Printf("The number I was thinking of was %d\n", myrand)
    }
}