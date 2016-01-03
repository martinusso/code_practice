package projecteuler

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	answer := problem1()
	if answer != answerProblem1 {
		t.Errorf("Expected %d got %d", answerProblem1, answer)
	}
}

func TestProblem2(t *testing.T) {
	answer := problem2()
	if answer != answerProblem2 {
		t.Errorf("Expected %d got %d", answerProblem2, answer)
	}
}

func TestProblem3(t *testing.T) {
	answer := problem3()
	if answer != answerProblem3 {
		t.Errorf("Expected %d got %d", answerProblem3, answer)
	}
}

func TestProblem4(t *testing.T) {
	answer := problem4()
	if answer != answerProblem4 {
		t.Errorf("Expected %d got %d", answerProblem4, answer)
	}
}
