package shapes

import "math"

type Shape interface {
	CalculateArea() float64
	CalculatePerimeter() float64
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) CalculatePerimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) CalculateArea() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) CalculatePerimeter() float64 {
	return math.Trunc(2 * math.Pi * c.Radius)
}

func (c Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}
