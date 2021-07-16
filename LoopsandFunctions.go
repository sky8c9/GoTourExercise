//Exercise: Loops and Functions
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	
	for {
		prevZ := z
		z -= (z*z - x) / (2*z)
		if math.Abs(z-prevZ) < 1e-8 {
			break;
		}
	}
	
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}