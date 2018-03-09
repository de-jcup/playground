package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	m2 := make(map[string]string)
	m2["key-a"] = "value-a"
	m2["key-b"] = "value-b"

	fmt.Println(m2)
	fmt.Printf("key-a=%s\n", m2["key-a"])
	fmt.Printf("key-b=%s\n", m2["key-b"])
	fmt.Printf("key-not-existing=%s\n", m2["key-not-existing"])
}
