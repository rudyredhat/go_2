// test the sqrt func from a values of csv file
package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

// open the file and check for the err
func TestMany(t *testing.T) {
	file, err := os.Open("sqrt_cases.csv")
	if err != nil {
		t.Fatalf("can't open cases file - %s", err)
	}
	defer file.Close()

	// create new reader and run infinite loop
	rdr := csv.NewReader(file)
	for {
		record, err := rdr.Read()
		// check for the EOF and break the loop
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("error reading cases file - %s", err)
		}

		// pass the first val
		val, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[0])
		}

		// pass the second val
		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("bad value - %s", record[1])
		}

		// run the subset and call the func and check for the error
		t.Run(fmt.Sprintf("%f", val), func(t *testing.T) {
			out, err := Sqrt(val)
			if err != nil {
				t.Fatal(err)
			}

			if !almostEqual(out, expected) {
				t.Fatalf("%f != %f", out, expected)
			}
		})
	}
}
