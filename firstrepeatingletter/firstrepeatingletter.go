package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(firstNonRepeating("sTress"))
	
}

func firstNonRepeating(word string) string {
	m := make(map[string]int)
	for i := 0; i < len(word); i++ {
		m[strings.ToLower(string(word[i]))] += 1	
	}
	for i := 0; i < len(word); i++ {
		if m[strings.ToLower(string(word[i]))] == 1 {
			return string(word[i])
		}
	}
	return ""
}