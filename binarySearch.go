package main

// Binary Search
func binarySearch(haystack []int, needle int) bool {
	// O(logN)
	lo := 0
   	hi := len(haystack) - 1

	for ok := true; ok; ok = (hi > lo) {
		m := lo + (hi-lo)/2
		v := haystack[m]
		if needle == v {
			return true
		} else if needle > v {
			lo = m + 1
		} else if needle < v {
			hi = m - 1
		}
	}
	return false
}
