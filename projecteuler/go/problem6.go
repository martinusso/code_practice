package projecteuler

// Problem 6
// Sum square difference
// The sum of the squares of the first ten natural numbers is,
// 1*2 + 2*2 + ... + 10*2 = 385
//
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)*2 = 55*2 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

const (
	answerProblem6 = 25164150
	oneHundred     = 100
)

func problem6() int64 {
	return squaresOfSum() - sumOfSquares()
}

func squaresOfSum() int64 {
	sumOfNumbers := int64((oneHundred * (oneHundred + 1)) / 2)
	return sumOfNumbers * sumOfNumbers
}

func sumOfSquares() int64 {
	var sum int64
	for i := int64(1); i <= oneHundred; i++ {
		sum += i * i
	}
	return sum
}
