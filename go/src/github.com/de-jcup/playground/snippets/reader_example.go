package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
/* my simple solution of https://tour.golang.org/methods/22 */
func (r MyReader) Read(b []byte) (n int, err error){
	n=1
	err=nil
	//for x:=range b{
	//	b[x]:='A'
	//}
	b[0]='A'
	return n,err
}

func main() {
	reader.Validate(MyReader{})
}
