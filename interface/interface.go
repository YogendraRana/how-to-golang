package main

import "fmt"

// interface 1
type PersonInterface interface {
	GetName() string
	GetAge() int
}

// struct 1
type PersonStruct struct {
	Name string
	Age  int
}

// implement interface 1
func (person PersonStruct) GetName() string {
	return person.Name
}

func (person PersonStruct) GetAge() int {
	return person.Age
}

// interface 2
type Shape interface {
	Area() float64
	Perimeter() float64
}

// struct 2
type Rectangle struct {
	Width  float64
	Height float64
}

// implement interface 2
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	var Alex PersonInterface = PersonStruct{
		Name: "Alex",
		Age:  20,
	}

	fmt.Printf("Name of alex is obviously %v\n", Alex.GetName())
	fmt.Printf("Age of alex is %v\n", Alex.GetAge())

	var shape Shape
	shape = Rectangle{Width: 4, Height: 5}

	area := shape.Area()
	perimeter := shape.Perimeter()

	fmt.Printf("Area of the rectangle is %v\n", area)
	fmt.Printf("Perimeter of the rectangle is %v\n", perimeter)

}
