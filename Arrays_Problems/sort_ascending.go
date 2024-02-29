// 14.Write a Program to sort the given array in ascending order
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 8, 4, 2, 3, 4, 5, 6, 7, 9}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}
