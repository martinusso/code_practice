package projecteuler

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	answer := 233168
	r := problem1()
	if r != answer {
		t.Errorf("Expected %d got %d", answer, r)
	}
}

func TestProblem2(t *testing.T) {
	answer := 4613732
	r := problem2()
	if r != answer {
		t.Errorf("Expected %d got %d", answer, r)
	}
}
