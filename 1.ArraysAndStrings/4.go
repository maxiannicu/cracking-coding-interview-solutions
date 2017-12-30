package main

import "fmt"
import "unicode"

// Runtime : O(n) , n - number of characters of input string
// Memory : O(1)
func isPalindromePermutation(a string) bool {
	vector := make([]int8, 1 << 8)
	oddElements := 0
	for _, el := range a {
		if unicode.IsLetter(el) {
			vector[el] += 1
		}
	}
	
	for _, el := range vector {
		if el % 2 == 1 {
			oddElements++
		}
	}

	return oddElements < 2
}

func main(){
	testInputs := []string{
		"tact coa",
		"coa tact",
		"abcba",
		"abccba",
		"abc ba",
	}
	
	for _, el := range testInputs {
		fmt.Printf("%v : %v\n", el, isPalindromePermutation(el))
	}
}