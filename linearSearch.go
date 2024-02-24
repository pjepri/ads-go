package main

// Linear Search
func linearSearch(haystack []int, needle int) bool {

	/// O(n)

	for i := 0; i < len(haystack); i++ {
		if needle == haystack[i] {
			return true

		}
	}
	return false
}