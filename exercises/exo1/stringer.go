package main

import (
	"fmt"
	"math"
)

//Shape is
type Shape interface {
	Area() float64
	fmt.Stringer
}

//Circle is
type Circle struct {
	r float64
}

//Rectangle is
type Rectangle struct {
	width  float64
	length float64
}

// Area is
func (r Rectangle) Area() float64 {
	return r.width * r.length
}

//Area is
func (c Circle) Area() float64 {
	return 3.14 * math.Pow(c.r, 2)
}
func (r Rectangle) String() string {
	return fmt.Sprint("Rectangle de largeur %d et de longueur %d", r.width, r.length)
}
func (c Circle) String() string {
	return fmt.Sprint("Cerlce de %d de rayon ", c.r)
}
func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
		fmt.Println("Et d'air : ", s.Area())
	}
}
