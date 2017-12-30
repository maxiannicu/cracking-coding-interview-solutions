package main

import "fmt"
import "strings"

// Runtime : O(a+b)
// Memory : O(a+b)
// finds columns and rows which contains 0 and then clears them
func isStringRotation(a,b string) bool {
	concat := fmt.Sprintf("%s%s",a,a)
	if strings.Contains(concat,b) {
		concat = strings.Replace(concat, b, "", 1)
		
		return strings.Compare(concat, a) == 0
	}
	return false
}

func main(){
	inputStrings := [][]string{
		{
			"waterbottle",
			"erbottlewat",
		},
		{
			"abcd",
			"bcda",
		},
		{
			"abcdef",
			"afcdeb",
		},
	}
	
	for _, pair := range inputStrings {
		fmt.Printf("%v\t%v\t%v\n",pair[0],pair[1],isStringRotation(pair[0],pair[1]))
	}
}