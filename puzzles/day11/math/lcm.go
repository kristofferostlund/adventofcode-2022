package math

import "fmt"

func LCMSlice(integers []int) int {
	if len(integers) < 2 {
		panic(fmt.Sprintf("no common multiple for %d numbers", len(integers)))
	}

	return lcm(integers[0], integers[1], integers[2:]...)
}

/*
Implementation of lcm and gcd are both copied from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
*/

// find Least Common Multiple (lcm) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// greatest common divisor (gcd) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
