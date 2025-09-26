package main

import "fmt"

type Shape struct {
	Rectangle
}

type Rectangle struct {
	length float64
	width float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64)  {
	r.length *= factor
	r.width *= factor
}

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle with a factor of 2 is", area)

	num := MyInt(-5)
	num1 := MyInt(10)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())
	fmt.Println(num1.welcomeMessage())

	// Shape
	s := Shape{
		Rectangle: Rectangle{length: 10, width: 9},
	}
	fmt.Println(s.Area())
	fmt.Println(s.Rectangle.Area())
}

type MyInt int
func (m MyInt) IsPositive() bool {
	return m > 0
}

func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

/*
-----------------------Output----------------------
Area of rectangle with width 9 and length 10 is 90
Area of rectangle with a factor of 2 is 360
false
true
Welcome to MyInt Type
Welcome to MyInt Type
90
90
-----------------------Output----------------------
*/