package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func matcher(pattern string, name string) {
	matches, err := filepath.Match(pattern, name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matches)
}

func matcher_glob(pattern string) {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(matches)
}

func main() {

	// just print current working dir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// test 1
	matcher_glob("./src/github.com/de-jcup/playground/files/sub*")
	matcher_glob("./src/github.com/de-jcup/playground/files/*sub*/*sub*/*.txt")
	matcher_glob("src/github.com/de-jcup/playground/files/*sub*/*sub*/*.txt")
	matcher_glob("src/github.com/de-jcup/playground/files")
	matcher_glob("src/github.com/de-jcup/*/*les")
	matcher_glob("*files*")

	// test 2
	matcher("*/src/github.com/de-jcup/*/*les", "./src/github.com/de-jcup/playground/xyz/files")
}
