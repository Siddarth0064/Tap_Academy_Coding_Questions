//2. Write a Program to print the number of non-zero elements.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 0, 4, 0, -5, 0, 6, 0, 7, 0, -8, 0}
	for _, val := range arr {
		if val > 0 || val < 0 {
			fmt.Print(val, " ")
		}
	}

}
