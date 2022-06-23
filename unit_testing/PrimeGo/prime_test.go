package Task3

import "testing"

type addTest struct {
    arg1 int 
		expected bool
}

var addTests = []addTest{
    addTest{11, true},
    addTest{13, false},
    addTest{99, false},
    addTest{111, false},
    
}


func TestCheckPrime(t *testing.T){

    for _, test := range addTests{
        if output := CheckPrime(test.arg1 ); output != test.expected {
            t.Errorf("Output %t not equal to expected %t", output, test.expected)
        }
    }
}
