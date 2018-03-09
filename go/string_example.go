package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/* this interface is "Stringer" which is defined in "fmt" */
/* see https://tour.golang.org/methods/17 for details */
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
