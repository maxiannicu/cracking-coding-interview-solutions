package main

import "fmt"
import "bytes"

// Runtime : O(n) , n - number of characters of input string
// Memory : O(1)
func compress(a string) string {
	var buffer bytes.Buffer
	writeChar := func(ch rune, repeats int){
		buffer.WriteString(fmt.Sprintf("%c%d", ch, repeats))
	}	
	
	lastChar := rune(0)
	charCount := 0
	length := len(a)
	
	for i, el := range a {
		if i == 0 {
			lastChar = el
			charCount = 1
		} else if i == length - 1 {
			if lastChar != el {
				writeChar(lastChar, charCount)
				writeChar(el, 1)
			} else {
				charCount++
				writeChar(lastChar, charCount)
			}
		} else {
			if lastChar == el {
				charCount++
			} else {
				writeChar(lastChar, charCount)
				charCount = 1
				lastChar = el
			}
		}
	}
	
	newResult := buffer.String()

	if length <= len(newResult) {
		return a
	}
	
	return newResult
}

func main(){
	inputValues := []string {
		"aabcccccaaa",
		"aabb",
		"abc",
	}
	
	for _,el := range inputValues {
		fmt.Printf("%v\t%v\n",el, compress(el))
	}
}