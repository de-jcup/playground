package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	fmt.Println("-- vertex map")
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println("-- string map")
	m2 := make(map[string]string)
	m2["key-a"] = "value-a"
	m2["key-b"] = "value-b"

	fmt.Println(m2)
	fmt.Printf("key-a=%s\n", m2["key-a"])
	fmt.Printf("key-b=%s\n", m2["key-b"])
	fmt.Printf("key-not-existing=%s\n", m2["key-not-existing"])
	
	
	fmt.Println("-- vertex map with literals")
	m3:= map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)
	
	mutateMaps()
}

func mutateMaps(){
	fmt.Println("-- mutate maps")
	m := make(map[string]int)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok = m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
