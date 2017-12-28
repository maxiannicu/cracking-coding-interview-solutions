package main

import "fmt"

// Runtime : O(a+b) , a,b - number of characters of input strings
// Memory : O(1)
func isPermutation(a,b string) bool {
	occurences := make([]int,1 << 8)
	for _, ch := range a {
		occurences[ch]++
	}
	for _, ch := range b {
		occurences[ch]--
	}
	for _, it := range occurences {
		if it != 0 {
			return false
		}
	}
	
	return true
}

func main(){
	inputsMap := map[string]string {
		"ab" : "abc",
		"adef" : "fdae",
		"hello" : "lehol",
	}

	for key,val := range inputsMap {
		fmt.Printf("%v & %v = %v\n",key,val,isPermutation(key,val))
	}
}
