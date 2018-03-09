package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	words := []string{"Hello", "World", "around", "at", "all"}
	var s []int = primes[1:4]
	var wordSlice []string = words[0:1]
	wordSlice2 := words[2:4]

	fmt.Println(s)
	fmt.Println(wordSlice)
	fmt.Println(wordSlice2)

}
