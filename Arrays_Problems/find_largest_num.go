// 4. Write a Program to find the largest number in an array.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 23, 45, 6, 7, 78}
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
