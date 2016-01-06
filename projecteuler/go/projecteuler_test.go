package projecteuler

import (
	"reflect"
	"testing"
)

type testPair struct {
	Name      string
	Got, Want int64
}

var testPairs = []testPair{
	{"problem 1", problem1(), answerProblem1},
	{"problem 2", problem2(), answerProblem2},
	{"problem 3", problem3(), answerProblem3},
	{"problem 4", problem4(), answerProblem4},
	{"problem 5", problem5(), answerProblem5},
	{"problem 6", problem6(), answerProblem6},
}

func TestProjectEuler(t *testing.T) {
	for _, pair := range testPairs {
		if !reflect.DeepEqual(pair.Got, pair.Want) {
			t.Errorf("%s: got %#v, want %#v", pair.Name, pair.Got, pair.Want)
		}
	}
}
