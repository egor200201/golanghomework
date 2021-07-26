package main

import (
	"fmt"
	"math"
)

type Shape interface {
	String() string
	Area() float64
	Perimeter() float64
}

func DescribeShape(s Shape) {
	fmt.Println(s)
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

type Circle struct {
	radius float64
}
type Rectangle struct {
	height float64
	width  float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Cirlce: radius %f", c.radius)
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle with height %f and width %f", r.height, r.width)
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {

	return 2 * (r.width + r.height)
}

func main() {
	c := Circle{radius: 6.434}
	r := Rectangle{
		height: 8.3874,
		width:  2.53334,
	}
	DescribeShape(c)
	DescribeShape(r)
}
