package projecteuler

// Problem 3
// Largest prime factor
//
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
//

const (
	answerProblem3 = 6857
)

func problem3() int {
	number := 600851475143

	for i := 2; i+i < number; i++ {
		for number%i == 0 {
			number /= i
		}
	}
	return number
}
