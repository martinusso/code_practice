package projecteuler

// Problem 1
// Multiples of 3 and 5
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

const (
	answerProblem1 = 233168
)

func problem1() int64 {
	sum := 0
	for n := 3; n < 1000; n++ {
		if isMultiplesOf3And5(n) {
			sum += n
		}
	}
	return int64(sum)
}

func isMultiplesOf3And5(number int) bool {
	return (number%3 == 0) || (number%5 == 0)
}
