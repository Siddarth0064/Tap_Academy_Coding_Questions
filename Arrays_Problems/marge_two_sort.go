//16.Write a Program to merge two arrays.

package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{6, 7, 8, 9, 10}
	result := append(arr1, arr2...)
	fmt.Println(result)
}
