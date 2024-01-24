package main

import "fmt"

// Define an interface named Shape
type Shape interface {
	Area() float64
}

// Implement the Shape interface for a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Shape interface for a circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Function that takes any Shape and calculates its area
func PrintArea(s Shape) {
	fmt.Printf("Area of the shape: %f\n", s.Area())
}

func main() {
	// Create instances of Rectangle and Circle
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}

	// Use the PrintArea function with different shapes
	PrintArea(rectangle)
	PrintArea(circle)
}
