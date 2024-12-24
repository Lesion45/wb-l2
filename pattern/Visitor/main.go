package main

import "fmt"

type Shape interface {
	Accept(visitor ShapeVisitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

type ShapeVisitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
}

type AreaCalculator struct{}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Square: %.2f\n", area)
}

func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("Square: %.2f\n", area)
}

type ShapeDrawer struct{}

func (s *ShapeDrawer) VisitCircle(c *Circle) {
	fmt.Printf("Draw with radius %.2f\n", c.Radius)
}

func (s *ShapeDrawer) VisitRectangle(r *Rectangle) {
	fmt.Printf("Rectangle with width %.2f and height %.2f\n", r.Width, r.Height)
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 5},
		&Rectangle{Width: 4, Height: 6},
	}

	areaCalculator := &AreaCalculator{}
	shapeDrawer := &ShapeDrawer{}

	for _, shape := range shapes {
		shape.Accept(areaCalculator)
	}

	for _, shape := range shapes {
		shape.Accept(shapeDrawer)
	}
}
