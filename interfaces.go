package main

import(
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perim() float64 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry){
	fmt.Print("Area:", g.area())
	fmt.Print("Perimeter:", g.perim())
}

func interfaceDemo(){
	r:= rectangle{width: 3, height: 4}
	c:= circle{radius: 5}
	fmt.Println("Rectangle:")
	measure(r)
	fmt.Println("Circle:")
	measure(c)
}

