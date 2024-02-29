//22.Write a Program to find the sum of diagonal elements of the matrix.
//                  1 2 3
//                  4 5 6       1+5+9=15
//                  7 8 9

package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	sumOfDiagonal := 0
	for i := 0; i < len(arr); i++ {
		sumOfDiagonal += arr[i][i]
	}
	fmt.Println(sumOfDiagonal)
}
