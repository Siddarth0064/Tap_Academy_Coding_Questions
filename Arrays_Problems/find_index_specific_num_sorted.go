//12.Write a Program to find the index of a specific number in an sorted array
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	findIndex := 5
	for i, val := range arr {
		if findIndex == val {
			fmt.Println(i, " ")
		}
	}
}
