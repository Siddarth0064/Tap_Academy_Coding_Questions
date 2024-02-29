//10.Write a Program to find the difference between the sum of even numbers & odd numbers.
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	sumOfOdd := 0
	sumOfEven := 0
	differenceOfEvenOdd := 0
	for _, val := range arr {
		if val%2 == 0 {
			sumOfEven += val
		} else {
			sumOfOdd += val
		}
	}
	differenceOfEvenOdd = sumOfEven - sumOfOdd
	fmt.Println(sumOfEven)
	fmt.Println(sumOfOdd)
	fmt.Println(differenceOfEvenOdd)
}
