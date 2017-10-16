package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    ret := 0
	first := 0
	second := 1
    tmp := 0
	return func() int {
        ret = first
        tmp = first + second
        first = second
        second = tmp
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
