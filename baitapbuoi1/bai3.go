package main

import (
	"fmt"
)

func prime(n uint32) string {
	s := []uint32{}
	var a uint32
	for a = 2; a <= n; a++ {
		var i uint32
		count := 0
		for i = 1; i < a; i++ {
			if a%i == 0 {
				count += 1
			}
		}
		if count == 1 {
			s = append(s, a)
		}
	}
	return fmt.Sprint(s)

}

func main() {
	fmt.Println(prime(100000))
}
