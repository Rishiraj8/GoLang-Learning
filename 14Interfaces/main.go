package main

import "fmt"

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    c := Circle{Radius: 5}
	fmt.Println("Area of Circle: ",c.Area())
    var s Shape = c
    fmt.Println(s.Area())
}
