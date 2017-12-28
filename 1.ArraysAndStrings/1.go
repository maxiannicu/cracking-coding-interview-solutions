package main

import "fmt"

// Runtime : O(n) , n - number of characters of input string
// Memory : O(1)
func hasUniqueCharacters(input string) bool {
	usageMatrix := make([]int, 1 << 8)

	for _, ch := range input {
		usageMatrix[ch]++
	}
	
	for _, ch := range usageMatrix {
		if ch > 1 {
			return false
		}
	}

	return true
}

func main(){
	testInputs := []string{
		"abcd",
		"aabcd",
		"run",
	}
	for _, input := range testInputs {
		fmt.Printf("%v : %v\n",input, hasUniqueCharacters(input))
	}
}
