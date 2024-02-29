// 19.Write a Program to find the index of the largest number
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6, 75, 8, 9}
	max := 0
	indexVal := 0
	for index, val := range arr {
		if val > max {
			max = val
			indexVal = index
		}
	}
	fmt.Printf("%v index of the %v max val", indexVal, max)
}
