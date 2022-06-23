package FibGo

import "testing"

type addFib struct {
	arg      int
	expected int
}

var addFibs = []addFib{
	addFib{4, 3},
	addFib{14, 377},
	addFib{0, 0},
}

func TestFib(t *testing.T) {

	for _, test := range addFibs {
		if output := fib(test.arg); output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
