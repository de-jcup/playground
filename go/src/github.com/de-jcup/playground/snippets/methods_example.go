package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}


/**
 *  A person in java style ...
 */
type Person struct{
	firstname, lastname string
}

/* this is wrong beause we got a copy...*/
func (p Person) setFirstNameWrong(s string){
	fmt.Printf("setFirstNameW1 called with =%s, p.firstname=%s\n",s,p.firstname)
	p.firstname=s
	fmt.Printf("setFirstNameW2 called with =%s, p.firstname=%s\n",s,p.firstname)
}

func (p *Person) setFirstName(s string){
	fmt.Printf("setFirstName1 called with =%s, p.firstname=%s\n",s,p.firstname)
	p.firstname=s
	fmt.Printf("setFirstName2 called with =%s, p.firstname=%s\n",s,p.firstname)
}

func (p Person) getFirstName() string{
	return p.firstname
}



func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
	
	p:= Person{"Albert","T."}
	fmt.Println(p)
	fmt.Println(p.getFirstName())
	p.firstname="xx"
	p.setFirstNameWrong("Wrong")
	p.setFirstName("Correct")
	fmt.Println(p)
	fmt.Println(p.getFirstName())
	
	
}
