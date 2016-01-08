package projecteuler

// Problem 7
// 10001st prime
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//
// What is the 10 001st prime number?

const (
	answerProblem7 = 104743
)

func problem7() int64 {
	var (
		count     int
		lastPrime int64
	)

	for i := 0; count < 10001; i++ {
		if isPrime(i) {
			lastPrime = int64(i)
			count++
		}
	}
	return lastPrime
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	}

	for i := 2; i <= (number / 2); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
