package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Perimieter() float64 {
	return 2 * (rect.Height + rect.Width)
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
