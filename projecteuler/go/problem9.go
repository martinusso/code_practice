package projecteuler

import "math"

// Problem 9
// Special Pythagorean triplet
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
// a**2 + b**2 = c**2
// For example, 3**2 + 4**2 = 9 + 16 = 25 = 5**2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

const (
	answerProblem9 = 31875000
)

func problem9() int64 {
	a := int64(0)
	for ; ; a++ {
		for b := a + 1; b < 1000; b++ {
			sum := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
			sqrt := math.Sqrt(sum)
			if sqrt != math.Trunc(sqrt) {
				continue
			}

			c := int64(sqrt)
			if (c > b) && (a+b+c == 1000) {
				return a * b * c
			}
		}
	}
}
