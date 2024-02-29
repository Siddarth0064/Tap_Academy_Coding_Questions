//11.Write a Program to find the second largest number without sorting.

package main

import "fmt"

func main() {
	arr := []int{1, 52, 3, 4, 95, 6, 7, 8, 29, 10}
	firstMax := 0
	secondMax := 0
	for _, val := range arr {
		if val > firstMax {
			firstMax, secondMax = val, firstMax
		} else if val > secondMax {
			secondMax = val
		}
	}
	fmt.Println(secondMax)
}
