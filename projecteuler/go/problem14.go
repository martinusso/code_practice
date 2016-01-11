package projecteuler

// Problem 14
// Longest Collatz sequence
// The following iterative sequence is defined for the set of positive integers:
//
// n → n/2 (n is even)
// n → 3n + 1 (n is odd)
//
// Using the rule above and starting with 13, we generate the following sequence:
//
// 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
// It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//
// Which starting number, under one million, produces the longest chain?
//
// NOTE: Once the chain starts the terms are allowed to go above one million.

const (
	answerProblem14 = 837799
)

func problem14() int64 {
	var longestChain, startingNumber int64

	for i := int64(2); i < 1000000; i++ {
		count := collatz(i)
		if count > longestChain {
			longestChain = count
			startingNumber = i
		}
	}
	return startingNumber
}

func collatz(n int64) int64 {
	count := int64(1)
	for ; n > 1; count++ {
		if isEven(n) {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return count
}
