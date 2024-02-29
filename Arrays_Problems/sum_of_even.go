//8. Write a Program to find the sum of all even numbers in the arrays

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sumOfEven := 0
	for _, val := range arr {
		if val%2 == 0 {
			sumOfEven += val
		}
	}
	fmt.Println(sumOfEven)
}
