package projecteuler

// Problem 5
// Smallest multiple
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

const (
	answerProblem5 = 232792560
)

func problem5() int64 {
	r := int64(2520)
	for ; ; r += 10 {
		if isDivisibleFrom1To20(r) {
			return r
		}
	}
}

func isDivisibleFrom1To20(number int64) bool {
	for i := int64(3); i <= 20; i++ {
		if number%i != 0 {
			return false
		}
	}
	return true
}
