package shapes

import "testing"

func assertCalculations(t *testing.T, result, expected float64) {
	t.Helper()

	if result != expected {
		t.Errorf("Expected %.12f but go %.12f", expected, result)
	}
}

func TestCalculatePerimeter(t *testing.T) {
	scenarios := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "Perimeter->Rectangle",
			shape:    Rectangle{Height: 10.0, Width: 10.0},
			expected: 40,
		},
		{
			name:     "Perimeter->Circle",
			shape:    Circle{Radius: 10},
			expected: 62.0,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := scenario.shape.CalculatePerimeter()

			assertCalculations(t, result, scenario.expected)
		})
	}
}

func TestCalculateArea(t *testing.T) {
	scenarios := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{
			name:     "Area->Rectangle",
			shape:    Rectangle{Height: 12.0, Width: 6.0},
			expected: 72.0,
		},
		{
			name:     "Area->Circle",
			shape:    Circle{Radius: 10},
			expected: 314.1592653589793,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := scenario.shape.CalculateArea()

			assertCalculations(t, result, scenario.expected)
		})
	}
}
