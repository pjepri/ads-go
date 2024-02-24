package main

import (
	"math"
)

// Two Crystal Balls problem
func twoCrystalBalls(breaks []bool) int {
	// O(sqrtN)
	
	jumpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpAmount

	for ; i < len(breaks); i = i + jumpAmount {
		if breaks[i] {
			break
		}
		//Breakpoint
	}

	i = i - jumpAmount

	// Find the closest breakpoint
	for j := 0; j <= jumpAmount && i < len(breaks); i,j = i + 1, j + 1 {
		if breaks[i] {
			return i 
		}
	}
	
	// Return -1 if there is no breakpoint found
	return -1

}
