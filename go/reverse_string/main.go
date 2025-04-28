package main

import (
	"fmt"
	"strings"
)

// Complete the solution so that it reverses the string passed into it.

// 'world'  =>  'dlrow'
// 'word'   =>  'drow'

func Solution(word string) string {

	stringArray := strings.Split(word, "")
	var newStringArray []string

	for i := range stringArray {
		newStringArray = append(newStringArray, stringArray[len(stringArray)-i-1])
	}

	res := strings.Join(newStringArray, "")

	return res
}

func main() {
	fmt.Println("Solution should return 'olleH':")
	fmt.Println(Solution("Hello"))
}

// https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go

func SolutionTwo(word string) string {
	var result = ""
	for _, c := range word {
		result = string(c) + result
	}

	return result
}
