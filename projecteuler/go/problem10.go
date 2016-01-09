package projecteuler

// Problem 10
// Summation of primes
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

import (
	"github.com/martinusso/go-lab/prime/sieveofatkin"
)

const (
	answerProblem10 = 142913828922
)

func problem10() int64 {
	var sum int64
	primes := sieveofatkin.GetPrimes(2000000)
	for _, v := range primes {
		sum += v
	}
	return sum
}
