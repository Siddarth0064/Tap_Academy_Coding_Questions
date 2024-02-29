//5. Write a Program to find the smallest number in an array

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 8, 9, 0, -3, -8, -66, 98}
	smallest := 0
	for _, val := range arr {
		if val < smallest {
			smallest = val
		}
	}
	fmt.Println(smallest)
}
