package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	
	return num
}

// Runtime : O(n) , n - number of characters of input string
// Memory : O(1)
func isOneAway(a,b string) bool {
	aLen := len(a)
	bLen := len(b)
	diff := abs(aLen-bLen)
	
	// if length diff is too much then we know for sure that there is more than a change
	if diff > 1 {
		return false
	}
	
	// if length is equal
	if bLen == aLen {
		changeFound := false
		for i := 0; i < aLen; i++ {
			if b[i] != a[i] {
				if changeFound {
					return false
				}
				changeFound = true
			}
		}
		
		return true
	}
	
	if bLen < aLen {
		a,b = b,a
		aLen,bLen = bLen,aLen
	}
	
	changeFound := false
	for i:=0 ; i < aLen; i++ {
		if changeFound && a[i] != b[i+1] {
			return false
		} else if !changeFound && a[i] != b[i] {
			changeFound = true
			
			if a[i] != b[i+1] {
				return false
			}
		}
	}
	
	return true
}

func main(){
	inputValues := map[string]string {
		"horse" : "horses",
		"abc" : "bca",
		"acd" : "abc",
		"horses" : "horse",
		"car" : "carter",
		"carer" : "cart",
	}
	
	for key,val := range inputValues {
		fmt.Printf("%v\t%v\t%v\n", key, val, isOneAway(key,val))
	}
}