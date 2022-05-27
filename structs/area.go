package structs

import "math"

// type Flags uint
// const (
// FlagUp Flags = 1 << iota // is up
// FlagBroadcast
// FlagLoopback
// FlagPointToPoint
// FlagMulticast
// )

// Приведенные далее инструкции объявляют структурный тип Employee и переменную dilbert, которая является экземпляром Employee:
// type Employee struct {
// ID int
// Name string
// Address string
// DoB time.Time
// Position string
// Salary int
// ManagerlD int
// }
// var dilbert Employee

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
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

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}
