package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	v2 := Vertex2{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// would fail: a = v

	// but:
	a = v2 //works because method v2.Abs uses vertex and not pointer...

	fmt.Println(a.Abs())
}

type MyFloat float64

/* ----------------- */
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
/* ----------------- */
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
/* ----------------- */

type Vertex2 struct {
	X, Y float64
}

/* this method uses a vertex and not the pointer (so cannot change and will always get a copy, but interesting for example */
func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
