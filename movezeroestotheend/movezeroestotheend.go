// move all the 0 elements in an array of numbers to the end of the array
package main

import "fmt"

func main() {
	// returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
	result := MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) 


	// [4, -3, -1, -2, -3, -5, 3, 3, 2, -1, 4, 1, 5, 3, -5, -4, -2, -5, 5, -2, 1, 3, 4, 3, 1, -3, 4, 2, 4, -1, -3, -3, -5, -2, -5, 5, 4, 1, 4, -2, 1, -3, -3, 3, 1, -4, -4, -4, 5, -1, 4, -5, -3, 4, 2, -5, -2, 1, -5, -2, 5, 2, 5, -3, -5, 2, -5, 3, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	// result := MoveZeros([]int{4, -3, -1, -2, -3, -5, 3, 3, 2, -1, 4, 1, 5, 3, -5, -4, -2, -5, 5, -2, 0, 3, 0, 3, 1, -3, 4, 0, 4, -1, -3, -3, -5, -2, -5, 5, 4, 1, 4, -2, 1, -3, -3, 3, 1, -4, -4, -4, 0, -1, 4, 0, -3, 4, 2, -5, -2, 0, -5, -2, 5, 2, 5, -3, -5, 2, -5, 0, 3, 3, 1, 0, -5, 5, 2, 4, 0, 1}) 
	fmt.Println(result)
}

func MoveZeros(arr []int) []int {	
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			if i == j {
				arr[j] = arr[i]
				j++
			} else {
				// it means that we encountered a 0, i is ahead, we can start replacing with 0s to have this an O(n) time and space
				arr[j] = arr[i]
				arr[i] = 0
				j++
			}			
		} 
	}
	return arr	
}