package main

import (
	"fmt"
	"math/rand"
)

func checknumber(a int) string {
	x := rand.Intn(100)
	switch {
	case a == x:
		return fmt.Sprint("Bạn đã đoán đúng")
	case a > x:
		return fmt.Sprint("Số bạn đoán lớn hơn X")
	default:
		return fmt.Sprint("Số bạn đoán nhỏ hơn X")
	}
}

func main() {
	fmt.Println(checknumber(2))
}
