package main

import (
	"io"
	"os"
	"fmt"
	"strings"
)
/* TODO: implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.*/
type rot13Reader struct {
	r io.Reader
}

/* my simple solution for https://tour.golang.org/methods/23 */
func (rot13 *rot13Reader) Read(b []byte) (n int, err error){
	n,err = rot13.r.Read(b)
	if (true){
		return n,err
	}
	for i:=0;i<n;i++{
		c :=b[i]
		fmt.Printf("i=%d,n=%d,c=%c",i,n,c)
		c+=13
		if (c>26){
			c=c-26
		}
		b[i]=c
	}
	return n,err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
