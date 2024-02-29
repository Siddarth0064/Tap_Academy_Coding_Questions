//25.Write a Program to find the sum of 2 matrices
//                  1 2 3     3 2 1          4 4 4
//                  4 5 6  +  1 4 5  =       5 9 11
//                  7 8 9     2 5 6          9 13 15

package main

import "fmt"

func main() {
	arr1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	arr2 := [][]int{{3, 2, 1}, {1, 4, 5}, {2, 5, 6}}
	fmt.Println(arr1)
	fmt.Println(arr2)

	result := make([][]int, len(arr1))
	for i := range result {
		result[i] = make([]int, len(arr1[i]))
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[i]); j++ {
			result[i][j] = arr1[i][j] + arr2[i][j]
		}
	}
	fmt.Println(result)
}
