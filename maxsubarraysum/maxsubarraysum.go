package main

import "fmt"

func main() {
	// maxSequence := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// should be 6: [4, -1, 2, 1]

	// maxSequence := []int{-2, -3, -1, -5, -4}
	maxSequence := []int{7, 4, 11, -11, 39, 36, 10, -6, 37, -10, -32, 44, -26, -34, 43, 43}

	fmt.Println(findMaxSum(maxSequence))
}

func findMaxSum(numbers []int) int {
	if len(numbers) == 0 || isArrayAllNegatives(numbers) {
		return 0
	}

	maxSum := 0
	for i := 0; i < len(numbers); i++ {
		for j := i; j < len(numbers); j++ {
			tempSum := sumSubArray(numbers, i, j) 
			if tempSum > maxSum {
				maxSum = tempSum
			}
		}
	}

	return maxSum
}

func isArrayAllNegatives(arr []int) bool {
	for _, v := range arr {
		if v >= 0 {
			return false
		}
	}
	return true
}

func sumSubArray(arr []int, start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += arr[i]
	}
	return sum
}