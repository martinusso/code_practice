package projecteuler

// Problem 17
// Number letter counts
// If the numbers 1 to 5 are written out in words: one, two, three, four, five,
// then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//
// If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//
//
// NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

import (
	"regexp"
	"strings"

	"github.com/martinusso/inflect"
)

const (
	answerProblem17 = 21124
)

func problem17() int64 {
	var size int
	for i := 1; i <= 1000; i++ {
		s, _ := inflect.IntoWords(float64(i))
		size += len(sanitize(s))
	}
	return int64(size)
}

func sanitize(value string) string {
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	safe := reg.ReplaceAllString(value, "")
	return strings.ToLower(strings.Trim(safe, ""))
}
