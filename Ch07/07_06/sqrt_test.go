// benchmarking imp ***
// use of go =  need to benchmarking it, and it must be fast , so that it is not slower
// when adding features and fixing bugs
// to run the  benchmark = go test -v bench .
// if interested in only the benchmark = go test -v bench . -run TTT
// need to profile it , so that we can see where its taking time
// can benchmark for a profile = go test -v bench . -run TTT -cpuprofile=prof.out
// then, go tool pprof prof.out
// list Sqrt // can see the runtime of the function
package sqrt

import (
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value - %f", val)
	}
}

// create a benchmarking
func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
