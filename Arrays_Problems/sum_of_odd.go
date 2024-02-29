// 9. Write a Program to find the sum of all odd number in the arrays

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sumOfOdd := 0
	for _, val := range arr {
		if val%2 == 1 {
			sumOfOdd += val
		}
	}
	fmt.Println(sumOfOdd)
}
