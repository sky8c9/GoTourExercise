//Errors
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	z := 1.0
	for {
		prevZ := z
		z -= (z*z - x) / (2*z)
		if math.Abs(z-prevZ) < 1e-8 {
			break;
		}
	}
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(14))
	fmt.Println(Sqrt(-2))
}