package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("Should properly add 2 + 2", func(t *testing.T) {
		expected := 4

		result := Add(2, 2)

		if result != expected {
			t.Errorf("Expected '%d' but got '%d'", expected, result)
		}
	})
}
