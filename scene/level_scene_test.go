package scene

import (
	"reflect"
	"testing"

	"github.com/telecoda/go-teletris/domain"
)

func TestNumberToDigits(t *testing.T) {

	testNum := 123

	digits := numberToDigits(testNum)

	if len(digits) != domain.MaxScoreDigits {
		t.Errorf("ScoreDigits wrong length Expected: %d got: %d", domain.MaxScoreDigits, len(digits))
	}

	expected := []int{0, 0, 0, 1, 2, 3}

	if !reflect.DeepEqual(digits, expected) {
		t.Errorf("ScoreDigits incorrect Expected: %d got: %d", expected, digits)
	}
}
