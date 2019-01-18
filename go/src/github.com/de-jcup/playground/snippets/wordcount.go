package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

/* my simple solution for https://tour.golang.org/moretypes/23 */

func WordCount(s string) map[string]int {
	result:= make(map[string]int)
	words:=strings.Fields(s)
	for i:= range words{
		word:=words[i]
		v, _ := result[word]
		v++
		result[word]=v
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
