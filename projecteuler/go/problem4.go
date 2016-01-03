package projecteuler

import "strconv"

// Problem 4
// Largest palindrome product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

const (
	answerProblem4 = 906609

	max3DigitNumber = 999
	min3DigitNumber = 100
)

func problem4() int {
	largestPalindrome := 0
	for i := max3DigitNumber; i >= min3DigitNumber; i-- {
		for j := i; j >= min3DigitNumber; j-- {
			product := i * j

			if product < largestPalindrome {
				break
			}

			if isPalindromic(product) {
				largestPalindrome = product
			}
		}
	}
	return largestPalindrome
}

func isPalindromic(number int) bool {
	n := strconv.Itoa(number)
	for i := 0; i < len(n)/2; i++ {
		if n[i] != n[len(n)-i-1] {
			return false
		}
	}
	return true
}
