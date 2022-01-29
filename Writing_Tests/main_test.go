package main

import "testing"

//Step2 - Table testing (better way)
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect 5", 50.0, 10.0, 5.0, false},
	{"expect fraction", -1.0, -777.0, 0.0012870013, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but didn't get one")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}

}

//Step 1 - inbox testing (bad way)
// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }

//Step1-a
// PS C:\Users\Ramin\go\src\learning-go\Writing_Tests> go test
// go: go.mod file not found in current directory or any parent directory; see 'go help modules'
// PS C:\Users\Ramin\go\src\learning-go\Writing_Tests> go mod init github.com/ramin/testprogram
// go: creating new go.mod: module github.com/ramin/testprogram
// go: to add module requirements and sums:
//         go mod tidy
// PS C:\Users\Ramin\go\src\learning-go\Writing_Tests> go test
// PASS
// ok      github.com/ramin/testprogram    0.276s

//Step1-b added TestBadDivide
// PS C:\Users\Ramin\go\src\learning-go\Writing_Tests> go test
// PASS
// ok      github.com/ramin/testprogram    0.268s

//After commenting step1

//Step2 - Table testing
// PS C:\Users\Ramin\go\src\learning-go\Writing_Tests> go test
// PASS
// ok      github.com/ramin/testprogram    0.290s
