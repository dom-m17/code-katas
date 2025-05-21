package main

import (
	"fmt"
	"strings"
)

// Complete the solution so that it splits the string into pairs of two characters.
// If the string contains an odd number of characters then it should replace the missing second character of the final pair with an underscore ('_').

// Examples:

// * 'abc' =>  ['ab', 'c_']
// * 'abcdef' => ['ab', 'cd', 'ef']

func Solution(str string) (res []string) {
	tmp := strings.Split(str, "")

	if len(tmp)%2 == 1 {
		tmp = append(tmp, "_")
	}

	for i, letter := range tmp {
		if i%2 == 0 {
			continue
		}
		res = append(res, fmt.Sprint(tmp[i-1], letter))
	}

	return res
}

func main() {
	fmt.Println(Solution("abc"))
	fmt.Println(Solution("abcdef"))
}
