package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func equations(a, b, c float64) string {
	delta := b*b - 4*a*c
	if a == 0 {
		return fmt.Sprint("a need#0")
	} else if delta == 0 {
		return fmt.Sprint(-(b / 2 * a))
	} else if delta > 0 {
		return fmt.Sprint(((-b + math.Sqrt(delta)) / 2), ((-b - math.Sqrt(delta)) / 2))
	}
	return fmt.Sprint((complex(-b, 0)+cmplx.Sqrt(complex(delta, 0)))/complex(2*a, 0), (complex(-b, 0)-cmplx.Sqrt(complex(delta, 0)))/complex(2*a, 0))
}

func main() {
	fmt.Println(equations(8, 4, 3))
}
