package basics

import (
	"fmt"
	"reflect"
)

func main() {
	//declare
	var x []int
	fmt.Println(reflect.ValueOf(x), reflect.ValueOf(x).Kind())
	//initilaise
	var Z = []int{1, 2, 3, 4, 5}
	fmt.Println(Z)
	//using make
	var Y = make([]int, 10, 20)
	fmt.Println(Y, len(Y), cap(Y))
	//using new
	var a = new([50]int)[0:10]
	fmt.Println(cap(a), a)

	// Add items using the append function
	var b = make([]int, 1, 10)
	fmt.Println(b)
	b = append(b, 20)
	fmt.Println(b)
	// Access slice items
	var c = []int{10, 20, 30, 40}
	fmt.Println(c[0])
	fmt.Println(c[0:3])
	// Change item values
	var d = []int{10, 20, 30, 40}
	fmt.Println(d)
	d[1] = 35
	fmt.Println(d)

	// Copy slice into another slice
	var e = []int{10, 20, 30, 40}
	var f = []int{50, 60, 70, 80}
	copy(e, f)
	fmt.Println("E: ", e)

	// Append a slice to an existing one
	var g = []int{10, 20, 30, 40}
	var h = []int{50, 60, 70, 80}

	g = append(g, h...)
	fmt.Println(g)
}
