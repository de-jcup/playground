package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	return guessSqrt(x, x/2, 0, 2000)
}

func guessSqrt(x, z float64, counter int, max int) float64 {
	/* guard close */
	if counter >= max {
		return z
	}

	if z*z == x {
		/* reached */
		return z
	}
	z -= (z*z - x) / (2 * z)
	return guessSqrt(x, z, counter+1, max)
}

func main() {
	fmt.Println(Sqrt(2))
}
