package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func isAnagramUsingReflect(s1, s2 string) bool {
	charFreq1 := make(map[rune] int)
	charFreq2 := make(map[rune] int)

	for _, rune := range s1 {
		charFreq1[rune] += 1 
	}

	for _, rune := range s2 {
		charFreq2[rune] += 1
	}
	fmt.Printf("Frequency map for %s %s", s1, s2)
	return reflect.DeepEqual(charFreq1, charFreq2)
}

func isAnagram(s1, s2 string) bool {
	charFreq := make(map[rune] int)

	for _, rune1 := range s1 {
		charFreq[rune1] += 1
	}

	for _, rune2 := range  s2 {
		_, ok := charFreq[rune2] 
		if !ok {
			return false
		}
		charFreq[rune2] -= 1
		if charFreq[rune2] == 0 {
			delete(charFreq, rune2)
		} 
	}

	return len(charFreq) == 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter string1: ")
	s1, _ := reader.ReadString('\n')
	fmt.Println("Enter string2: ")
	s2, _ := reader.ReadString('\n')

	fmt.Println("Is anagram (using reflect)? - %v", isAnagramUsingReflect(s1, s2))
	fmt.Println("Is anagram? - %v", isAnagram(s1, s2))
}