package main

import (
	"fmt"
	"math"
)

func main() {

	// must return 2
	// cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200}); 

	fmt.Println(cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}))
}

func cakes(recipe, ingredients map[string]int) int {
	min := math.MaxInt32
	for k := range recipe {
		if val, ok := recipe[k]; ok {
			tempMin := ingredients[k] / val
			if tempMin < min {
				min = tempMin
			}
		} else {
			return 0
		}
	}	
	return min
}
