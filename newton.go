package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
    y := x- (((x * x) - x) / 2 * x )
	return y
}

func main() {
	fmt.Println(Sqrt(Sqrt(Sqrt(3))))
}

