package structs

import "math"

// Shape interface
// GO的这个接口奇奇怪怪的，接口的实现类不需要显示的表明自己要去实现哪个接口
// 这种interface resolution是隐式的
// 当我将一个实现者当作接口传入的时候，才会去检查这个实现者有没有实现所有的接口定义
type Shape interface {
	Area() float64

	// 测试用，如果有未实现的接口，编译会报错
	// cannot use c (type Circle) as type Shape in argument to checkArea:
	// Circle does not implement Shape (missing Test method)

	// Test() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

// Area func for Rectable, width * height
func (r Rectangle) Area() (area float64) {
	return r.Height * r.Width
}

// Area func for Circle, pi * radius^2
func (c Circle) Area() (area float64) {
	return math.Pi * c.Radius * c.Radius
}

// Area func for Triangle, base * height * 1/2
func (t Triangle) Area() (area float64) {
	return t.Base * t.Height * 0.5
}
