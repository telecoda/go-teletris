package scene

import (
	"fmt"
	"testing"
)

func TestNumberToDigits(t *testing.T) {

	testNum := 123

	digits := numberToDigits(testNum)

	if len(digits) != 3 {
		t.Errorf("Error converting number to digits. Expected: %d got: %d", 3, len(digits))
	}

	fmt.Printf("Result: %#v\n", digits)
}
