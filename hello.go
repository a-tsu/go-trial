package main

import (
    "fmt"
    "math"
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Printf("Hello, 世界\n")

    fmt.Println(math.Pi)

    fmt.Println(add(42, 13))

    a, b := swap("hello", "world")
    fmt.Printf(b)
    fmt.Printf(a)
    fmt.Printf("\n")

    fmt.Println(split(122))
}
