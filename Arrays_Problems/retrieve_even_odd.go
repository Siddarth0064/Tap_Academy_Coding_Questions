//3. Write a Program to print the number of even elements & odd elements.

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	even := []int{}
	odd := []int{}
	for _, val := range arr {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	fmt.Println(even)
	fmt.Println(odd)
}
