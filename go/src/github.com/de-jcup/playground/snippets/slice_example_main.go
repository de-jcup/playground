package main

import "fmt"

func main() {
	fmt.Println("----nil:")
	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Println("snil=nil!")
	}

	fmt.Println("----ints:")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	printSliceInt(s)

	fmt.Println("----word:")
	words := []string{"Hello", "World", "around", "at", "all"}
	var wordSlice []string = words[0:1]
	wordSlice2 := words[2:4]

	fmt.Println(words)
	fmt.Println(wordSlice)
	fmt.Println(wordSlice2)

	/* change the slice ... and so words...*/
	wordSlice2[0] = "no longer around"
	fmt.Println(words)

	fmt.Println(">append")
	words = append(words, "and", "back")
	fmt.Println(words)

	fmt.Println("----structs:")

	/* structures as arrays/slices*/
	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	fmt.Println("--range over slices")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}

}

func printSliceInt(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
