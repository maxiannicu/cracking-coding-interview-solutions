package main

import "fmt"

// Runtime : O(n^2)
// Memory : O(n^2)
// creates new matrix with same size and copies data to the computed position
func invert(matrix [][]int) [][]int {
	n := len(matrix)
	result := make([][]int, n)
	for i:=0; i < n; i++ {
		result[i] = make([]int,n)
	}
	
	for i:=0; i < n; i++ {
		for e:=0; e < n; e++ {
			result[e][n-i-1] = matrix[i][e]
		}
	}
	
	return result
}

// Runtime : O(n^2)
// Memory : O(1)
// use 4 pivots (in each corner) and replace once a time. otherwise you will lose a value. this methods swaps values from edges to center
func invertInPlace(matrix [][]int) [][]int {
	n := len(matrix)
	
	for i:=0; i < n/2; i++ {
		for e:=i; e < n-i-1; e++ {
			matrix[e][n-i-1],matrix[n-i-1][n-e-1],matrix[n-e-1][i],matrix[i][e] = matrix[i][e],matrix[e][n-i-1],matrix[n-i-1][n-e-1],matrix[n-e-1][i]
		}
	}
	
	return matrix
}

func main(){
	matrix := [][]int{
		{ 1,2,3,4 },
		{ 5,6,7,8 },
		{ 9,10,11,12 },
		{ 13,14,15,16 },
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
	printMatrix(invert(matrix))
	fmt.Println("After invert in place")
	printMatrix(invertInPlace(matrix))
}