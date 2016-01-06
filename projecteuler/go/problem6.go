package projecteuler

import "math"

// Problem 6
//
// Sum square difference
// The sum of the squares of the first ten natural numbers is,
//
// 1*2 + 2*2 + ... + 10*2 = 385
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)*2 = 55*2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

const (
	answerProblem6 = 25164150
)

func problem6() int64 {
	var (
		sumOfSquares, squaresOfSum float64
		sumOfNumbers               int
	)
	for i := 1; i <= 100; i++ {
		sumOfSquares += (math.Pow(float64(i), 2))
		sumOfNumbers += i
	}
	squaresOfSum = math.Pow(float64(sumOfNumbers), 2)
	return int64(squaresOfSum - sumOfSquares)
}
