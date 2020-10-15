package main

import {
	"fmt"
	"math"
}
type Shape interface {
	Area() float64
	fmt.Stringer
}
type Circle struct {
	R float64
}
type Rectangle struct {
	width float64
	length float64
}
func (r Rectangle) Area() float64 {
	return r.width * r.length
}
func (c Circle) Area() float64 {
	return 3.14 * math.Pow(c.r,2)
}
func (r Rectangle) String() string {
	return fmt.Sprint("Rectangle de largeur %f et de longueur %f.", r.width, r.length)
}
func (c Circle) String() string {
	return fmt.Sprint("Cerlce de %f de rayon.", c.R)
}
func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		fmt.Println("Et d'air : "s.Area())
	}
}