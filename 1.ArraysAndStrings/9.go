package main

import "fmt"
import "strings"

// Runtime : O(a+b)
// Memory : O(a+b)
// Example :
// a = abcd
// b = bcda
// 1 - concatenate a => abcdabcd
// 2 - check if result contains b
// 3 - if no, return false. otherwise go to 5
// 4 - remove b from concatination => abcd
// 5 - if string after removal is equal to a , then it is rotation. otherwise no ! ...
//  	we can have a situation (a = abcd, b=bcd) which passes all steps but it's not a rotation. so this step assumes correct solution
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