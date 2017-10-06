//Sources: https://adriann.github.io/programming_problems.html
//https://coderwall.com/p/a6ghja/checks-for-palindrome-strings-or-numbers-in-go-lang
package main

//Package fmt implements formatted I/O
import (
	"fmt"
	"strings"
)

func main() {
	var palindrome string
	//Here we print out a question and then when the word
	//Is entered it is then changed to lowercase.
	fmt.Println("Enter a word below:")
	fmt.Scanf("%s\n", &palindrome)
	palindrome = strings.ToLower(palindrome)
	fmt.Println(isP(palindrome))
}

//This function will be used to test if the word 
//entered is actually a palindrome
func isP(s string) string {
	mid := len(s) / 2
	last := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[last-i] {
			return "No this word is not a Palindrome."
		}
 	}
	return "Yes the word you entered is a Palindrome"
}