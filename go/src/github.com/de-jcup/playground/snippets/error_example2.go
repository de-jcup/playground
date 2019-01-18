package main

import (
	"fmt"
)
/* my simple solution for https://tour.golang.org/methods/20 */
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	// leads to stack overflow: return fmt.Sprint(e)
	// -> reason: Sprintf knows error->try to call it again -->loop
	// leads NOT to stack overflow: return fmt.Sprint(float64(e))
	return fmt.Sprintf("Negative sqrt not allowed in this space. number was %f",e)
}

func Sqrt(x float64) (float64, error) {
	if (x<0){
		return 0,ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
