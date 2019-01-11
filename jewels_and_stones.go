package main

import (
	"fmt"
)

func main() {
	n := numJewelsInStones("ebd", "bbb")
	fmt.Println(n)
}

func numJewelsInStones(J string, S string) int {
	counter := 0
	for _, char := range J {
		for _, ch := range S {
			if char == ch {
				counter++
			}
		}
	}

	return counter
}
