//18.Write a Program to store the square of all the numbers in the resultant array

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	squareOfNum := []int{}
	for _, val := range arr {
		squareOfNum = append(squareOfNum, val*2)
	}
	fmt.Println(squareOfNum)
}
