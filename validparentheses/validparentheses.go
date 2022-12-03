package main

import "fmt"

func main() {
	// parens := "(())((()())())";
	parens := ")))(((";
	fmt.Println(isOrdered(parens))
}

func isOrdered(parens string) bool {
	center := 0
	for _, v := range parens {
		if v == '(' {
			center++
		} else {
			center--
		}
		if center < 0 {
			return false
		}
	}
	return center == 0
}