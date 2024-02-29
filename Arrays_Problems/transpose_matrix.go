//24.Write a Program to print the transpose of a matrix.
//                  1 2 3                        1 4 7
//                  4 5 6      transpose is =    2 5 8
//                  7 8 9                        3 6 9

package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr)
	transposeMatris := [][]int{}
	for i := 0; i < len(arr[0]); i++ {
		row := []int{}
		for j := 0; j < len(arr); j++ {
			row = append(row, arr[j][i])
		}
		transposeMatris = append(transposeMatris, row)
	}
	fmt.Println(transposeMatris)
}
