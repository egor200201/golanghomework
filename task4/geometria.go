package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	String() string
	Area() (float64, error)
	Perimeter() (float64, error)
}

func DescribeShape(s Shape) {
	var (
		v float64
		e error
	)
	v, e = s.Area()
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(s)
		fmt.Printf("Area: %.2f\n", v)
		v, _ = s.Perimeter()
		fmt.Printf("Perimeter: %.2f\n", v)
	}
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

func (c Circle) Area() (float64, error) {
	if c.radius <= 0 {
		return 0, errors.New("wrong values for circle")
	}
	return c.radius * c.radius * math.Pi, nil
}

func (c Circle) Perimeter() (float64, error) {
	if c.radius <= 0 {
		return 0, errors.New("wrong values for circle")
	}
	return 2 * math.Pi * c.radius, nil
}

func (r Rectangle) String() string {

	return fmt.Sprintf("Rectangle with height %f and width %f", r.height, r.width)
}

func (r Rectangle) Area() (float64, error) {
	if (r.height <= 0) || (r.width <= 0) {
		return 0, errors.New("wrong values for rectangle")
	}
	return r.width * r.height, nil
}

func (r Rectangle) Perimeter() (float64, error) {
	if (r.height <= 0) || (r.width <= 0) {
		return 0, errors.New("wrong values for rectangle")
	}
	return 2 * (r.width + r.height), nil
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
