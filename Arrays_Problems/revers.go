//21.Write a Program to reverse an array.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 6, 75, 8, 9}
	fmt.Println(arr)
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	fmt.Println(arr)
}
