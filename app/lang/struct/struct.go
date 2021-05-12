package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) circleArea() float64 {
	return math.Pi * c.radius * c.radius
}

func circleArea(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (rectangle *Rectangle) area() float64 {
	l := distance(rectangle.x1, rectangle.y1, rectangle.x1, rectangle.y2)
	w := distance(rectangle.x1, rectangle.y1, rectangle.x2, rectangle.y1)
	return l * w
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func main() {
	var c = Circle{x: 0, y: 0, radius: 5}
	var circle2 = Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.radius)
	fmt.Println(c.x, c.y, c.radius)
	circle2.x = 10
	circle2.y = 5
	fmt.Println(circle2.x, circle2.y, circle2.radius)
	fmt.Println(circleArea(circle2))
	fmt.Println(circle2.circleArea())
	r := Rectangle{
		x1: 0,
		y1: 0,
		x2: 10,
		y2: 10,
	}

	fmt.Println(r.area())

	person := Person{
		"Mike",
	}

	a := Android{
		person,
		"Google pixel",
	}

	a.Person.Talk()
	fmt.Println(a.Model)
	a.Person.Talk()
	fmt.Println(totalArea(&circle2, &r))
}
