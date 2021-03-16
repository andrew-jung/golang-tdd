package perimeter

import "math"

// Rectangle and Circle satisfies this interface since they implement an Area() func.
type Shape interface {
	Area() float64
}

// Struct is a named collection of fields where you can store data.
// Kinda like python enums
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Can't override same function definition for different paramters (ie Circle and Rectangle)
// 1. Create a different package for each Shape/Struct method
// 2. Define methods on our types instead, funcs have a receiver name that can take in our structs
// func (receiverName ReceiverType) funcName(args) returnValType {}, receiver is kind of like `this` in other languages

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }
