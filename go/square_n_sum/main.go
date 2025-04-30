package main

import "fmt"

// Complete the square sum function so that it squares each number passed into it and then sums the results together.

// For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9

func SquareSum(numbers []int) (result int) {
	for i := range numbers {
		result += numbers[i] * numbers[i]
	}

	return
}

func main() {
	fmt.Println(SquareSum([]int{1, 2}))       // Should equal 5
	fmt.Println(SquareSum([]int{0, 3, 4, 5})) // Should equal 50
	fmt.Println(SquareSum([]int{}))           // Should equal 0
}

// https://www.codewars.com/kata/515e271a311df0350d00000f/go

// No need really to use the index as the data is not large, should just use the value instead
// for _, num := range numbers {
// 		result += num * num
// }

// Bit more readable imo
