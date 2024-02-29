//23.Write a Program to find the largest number in the matrix.
//                  1 2 3
//                  4 5 6       largest num is 9
//                  7 8 9

package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	largestNum := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] > largestNum {
				largestNum = arr[i][j]
			}
		}
	}
	fmt.Println(largestNum)
}
