package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	t.Run("Should properly add 2 + 2", func(t *testing.T) {
		expected := 4

		result := Add(2, 2)

		if result != expected {
			t.Errorf("Expected '%d' but got '%d'", expected, result)
		}
	})
}
