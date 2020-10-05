/*
	1) come with build in testing framework
	2) test file in a go ends with = sqrt_test_file_name.go
	3) go run -v test_name -v // to eexcute single test
*/
package sqrt

import (
	"fmt"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		// Fatalf will stop the execution right here
		t.Fatalf("error in calculation - %s", err)
	}

	// compare the val
	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

type testCase struct {
	value    float64
	expected float64
}

// with multiple test cases
func TestMany(t *testing.T) {
	testCases := []testCase{
		{0.0, 0.0},
		{2.0, 1.414214},
		{9.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f", tc.value), func(t *testing.T) {
			out, err := Sqrt(tc.value)
			if err != nil {
				t.Fatal("error")
			}

			if !almostEqual(out, tc.expected) {
				t.Fatalf("%f != %f", out, tc.expected)
			}
		})
	}
}
