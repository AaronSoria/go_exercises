package main
import "testing"


var tests = []struct {
    name string
    divident float64
    divisor float64
    expected float64
    isErr bool
}{
    {"valid-data", 100.0, 10.0, 10.0, false}
    {"invalid-data", 100.0, 0.0, 0.0, true}
}

func testDivision(t *testing.T){
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected error but did not get one")
			}
		} else {
			t.Error("did not expect an error but got one", err.Error())
		}

		if got != tt.expected {
			t.Error("expectet %f but got %f", tt.expected, got)
		}
	}
}



func TestDivide(t *testing.T){
	 _, err:= divide(10.1, 1.0)
	 if err != nil {
		t.Error("Got an error")
	 }
}

// go test
// go test -v