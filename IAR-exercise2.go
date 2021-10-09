package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	a, b, d float32
}

func (c Circle) Area() float32 {
	return c.a * c.b * c.d
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	c := Circle{1, 2, 3}

	shaper := []Shaper{r, q, c}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shaper {
		fmt.Println("Shape details: ", shaper[n])
		fmt.Println("Area of this shape is: ", shaper[n].Area())
	}
}
