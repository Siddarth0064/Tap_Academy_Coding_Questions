//7. Write a Program to find the product of all numbers in the arrays

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	product := 1
	for _, val := range arr {
		product *= val
	}
	fmt.Println(product)
}
