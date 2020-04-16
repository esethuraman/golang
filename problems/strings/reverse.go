package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string...")
	s, _ := reader.ReadString('\n')
	fmt.Println(s)
	fmt.Println("Reversed value : ", reverse(s))
}
