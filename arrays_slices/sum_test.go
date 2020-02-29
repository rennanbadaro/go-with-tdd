package arraySlices

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
