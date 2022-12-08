package main

import (
	"fmt"

)

func main() {
	n := 16
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}

	E := n + 2
	S := n + 2
	W := -3
	N := -1
	X := 0
	Y := 0

	direction := "E"

	a[X][Y] = 1
	done := false
	for !done {
		X, Y = getNextCoordinatesAndDirection(X, Y, direction)
		switch direction {
		case "E":
			if Y < E - 2 && belowCellIsNotOne(a, X, Y) {
				a[X][Y] = 1
			} else if Y == E - 2 {
				direction = "S"
				E = E - 2
				Y = Y - 1
			} else {
				done = true
			}
		case "S":
			if X < S - 2 {
				a[X][Y] = 1
			} else if X == S - 2 {
				direction = "W"
				S = S - 2
				X = X - 1
			} else {
				done = true
			}
		case "W":
			if Y > W + 2 && aboveCellIsNotOne(a, X, Y) {
				a[X][Y] = 1
			} else if Y == W + 2 {
				direction = "N"
				W = W + 2
				Y = Y + 1
			} else {
				done = true
			}
		case "N":
			if X > N + 2 {
				a[X][Y] = 1
			} else if X == N + 2 {
				direction = "E"
				N = N + 2
				X = X + 1
			} else {
				done = true
			}
		}
	}
}

func aboveCellIsNotOne(a[][]int, X, Y int) bool {
	if X-1 < 0 {
		return true
	}
	return a[X-1][Y] == 0
}

func belowCellIsNotOne(a[][]int, X, Y int) bool {
	if X+1 > len(a) {
		return true
	}
	return a[X+1][Y] == 0
}

func getNextCoordinatesAndDirection(X, Y int, direction string) (int, int) {
	switch direction {
	case "E":
		return X, Y+1
	case "S":
		return X+1, Y
	case "W":
		return X, Y-1
	case "N":
		return X-1, Y		
	}
	return X, Y
}

func printMatrix(a [][]int) {
	for i := 0; i<len(a); i++ {
		for j := 0; j<len(a[i]); j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}