package main

import "fmt"

// Runtime : O(n*m)
// Memory : O(n+m)
// finds columns and rows which contains 0 and then clears them
func zero(matrix [][]int) [][]int {
	cc := len(matrix[0])
	rc := len(matrix)
	rows := make([]bool,rc)
	columns := make([]bool,cc)
	
	for r, row := range matrix {
		for c, el := range row {
			if el == 0 {
				rows[r] = true
				columns[c] = true
			}
		}
	}
	
	for i, r := range rows {
		if r {
			for e:=0; e < cc; e++ {
				matrix[i][e] = 0
			}
		}
	}
	
	for i, c := range columns {
		if c {
			for e:=0; e < rc; e++ {
				matrix[e][i] = 0
			}
		}
	}
	
	return matrix
}

func main(){
	matrix := [][]int{
		{ 1,0,3,5 },
		{ 5,6,1,5 },
		{ 9,10,11,5 },
	}
	
	printMatrix := func(matrix [][]int){
		for _, row := range matrix {
			for _, el := range row {
				fmt.Printf("%d ",el)
			}
			fmt.Println()
		}
	}
	
	fmt.Println("Before")
	printMatrix(matrix)
	fmt.Println("After")
	printMatrix(zero(matrix))
}