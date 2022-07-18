package main

import "math"

func calTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLen := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= 0 {
			start = lastI + 1
		}
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLen
}
func main() {

}
