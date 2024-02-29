//20.Write a Program to find the index of the smallest number.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6, 75, -8, 9}
	smallest := arr[0]
	indexVal := 0
	for index, val := range arr {
		if val < smallest {
			smallest = val
			indexVal = index
		}
	}
	fmt.Printf("%v index of the %v max val", indexVal, smallest)
}
