package main

import "fmt"
import "unicode"

// Runtime : O(n) , n - number of characters of input string
// Memory : O(1)
func urlfy(input []rune){
	right := len(input) - 1
	i := right
	for ; i >= 0 && !unicode.IsLetter(input[i]); i-- {
		// iterate right to left and find first character
	}
	for ; i >= 0; i-- {
		switch {
			case unicode.IsLetter(input[i]) : 
				input[right] = input[i]
				right--
			case input[i] == ' ' :
				input[right-2] = '%'
				input[right-1] = '2'
				input[right] = '0' 
				right -= 3
		}
	}
}

func main(){
	inputs := []string{
		"Mr John Smith    ",
		"Mr John  ",
	}
	for _,it := range inputs {
		conv := []rune(it)
		urlfy(conv)
		fmt.Printf("%v => %s\n", it, string(conv)) 
	}
}
