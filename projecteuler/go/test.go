package projecteuler

import (
	"testing"
)

func Test_Problem1(t *testing.T) {
	answer := 233168
	r := problem1()
	if r != answer {
		t.Errorf("Expected %d got %d", answer, r)
	}
}
