package arraySlices

import (
	"reflect"
	"testing"
)

func assertCollectionSum(t *testing.T, expected, result []int) {
	t.Helper()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v but got %v", expected, result)

	}
}

func TestSum(t *testing.T) {
	t.Run("Should properly sum a collection of N numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("Expected %d but got %d", expected, result)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Should properly sum a collection of slices returning a new slice with the sum of each of them", func(t *testing.T) {
		result := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		expected := []int{6, 15}

		assertCollectionSum(t, expected, result)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Should propely sum the tail of a given collection of slices", func(t *testing.T) {
		result := SumAllTails([]int{2, 3, 4, 5}, []int{0, 2})
		expected := []int{12, 2}

		assertCollectionSum(t, expected, result)
	})

	t.Run("Should propely sum the tail of a given collection of slices even if one of them is an empty slice", func(t *testing.T) {
		result := SumAllTails([]int{2, 3, 4, 5}, []int{})
		expected := []int{12, 0}

		assertCollectionSum(t, expected, result)
	})
}
