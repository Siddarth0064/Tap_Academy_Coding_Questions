//12.Write a Program to find the index of a specific number in an unsorted array

package main

import "fmt"

func main() {
	arr := []int{1, 8, 4, 2, 3, 4, 5, 6, 7, 9}
	findIndex := 7
	for i, val := range arr {
		if findIndex == val {
			fmt.Println(i, " ")
		}
	}
}
