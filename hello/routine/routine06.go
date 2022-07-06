package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	type T struct{ a, b, c int }
	var x = T{1, 2, 3}
	var y = T{4, 5, 6}
	var z = T{7, 8, 9}
	var v atomic.Value
	v.Store(x)
	fmt.Println(v) // {{1 2 3}}
	old := v.Swap(y)
	fmt.Println("old:", old)
	fmt.Println("v:", v)            // {{4 5 6}}
	fmt.Println("old.(T)", old.(T)) // {1 2 3}
	swapped := v.CompareAndSwap(x, z)
	fmt.Println(swapped, v) // false {{4 5 6}}
	swapped = v.CompareAndSwap(y, z)
	fmt.Println(swapped, v) // true {{7 8 9}}
}
