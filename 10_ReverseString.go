//Sources: https://adriann.github.io/programming_problems.html
//https://github.com/golang/example/blob/master/stringutil/reverse.go
//https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go

package main

//Package fmt implements formatted I/O
import (
	"fmt"
)

func main() { 
	//Enter a sentence below
    input := "SKY NEWS: Tory MPs overwhelmingly back May to remain PM" 
	//Get the Unicode code 
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input { 
			rune[n] = r
			n++
	} 
	rune = rune[0:n]
	//Reverse the string
	for i := 0; i < n/2; i++ { 
			rune[i], rune[n-1-i] = rune[n-1-i], rune[i] 
	}
	output := string(rune)
	fmt.Println(output)
}