// 1. Write a Program to print the number of elements in an array
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, val := range arr {
		fmt.Print(val, " ")
	}
}
