package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	runes := [] rune(s)
	// truncating the last new line char that was appended as part of reader
	runes = runes[0: len(runes)-1]
	
	fmt.Println(runes)
	
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			fmt.Println(runes[i], runes[j])
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string: ")
	s, _ := reader.ReadString('\n')
	fmt.Printf("Is %s a palindrome? - %v\n", s, isPalindrome(s))
}
