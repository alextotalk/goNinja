package main

import (
	"fmt"
	"math"
)

//type Shape interface {
//	Area() float32
//}
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}
type ShapeWithArea interface {
	Area() float32
}
type ShapeWithPerimeter interface {
	Perimeter() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Perimeter() float32 {
	return s.sideLength * 4
}

func (s Square) Area() float32 {
	return s.sideLength * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Perimeter() float32 {
	return 2 * c.radius * math.Pi
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}
	fmt.Println("Perimeter circle ", circle.Perimeter())
	printArea(square)
	printArea(circle)

	printInterface(square)
	printInterface(circle)
	printInterface("wqwq")
	printInterface(222)
	printInterface(true)
}
func printArea(shape Shape) {
	fmt.Println(shape.Area())
}
func printInterface(i interface{}) {
	str, ok := i.(string)
	if !ok {
		fmt.Println("interface not string")
		return
	}
	fmt.Println(str)

	//switch value := i.(type){
	//case int:
	//	fmt.Println("It is int", value)
	//case bool:
	//	fmt.Println("It is bool", value)
	//default:
	//	fmt.Println("It is unknown type" )
	//}
	//fmt.Printf("%+v\n",i)
}
