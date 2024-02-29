//6. Write a Program to find the sum of all elements in an array.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sum := 0
	for _, val := range arr {
		sum += val
	}
	fmt.Println(sum)
}
