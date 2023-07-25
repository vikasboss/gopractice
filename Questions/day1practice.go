package main

import (
	"fmt"
	"reflect"
)

func doOperationOnRectange(length, breadth int) (perimeter, area int) {
	area = length * breadth
	perimeter = 2 * (length + breadth)
	return perimeter, area
}

// anonymous function
var (
	xfunc = func(x int, y int) {
		fmt.Println(x, y)
	}
)

const PI = 3.141

// closure function
var (
	yfunc = func(x int, y int) {
		fmt.Println(x, y, PI)
	}
)

// variadic function
func variadicExample1(x ...string) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

func variadicExample2(a string, y ...interface{}) {
	for _, v := range y {
		fmt.Println(reflect.TypeOf(v), reflect.ValueOf(v))
	}
}

// defereed function
func def() {
	//In case of error this acts like fallback mechanism
	defer variadicExample1("x", "y", "Z")
	fmt.Println("Hello Wolrd")
	return
}

// Higher Order Function
func HighOrderFunctions() func() {
	return def
}
func main() {
	//declaration
	var x int
	//initaialisation
	x = 10
	fmt.Println(x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
	xy := 10
	fmt.Println(xy)

	var (
		xiwjiwjd  int
		xiwjiwjd2 int = 20
		stroigg       = "owksowksowks"
	)
	xiwjiwjd = 100
	fmt.Println(xiwjiwjd, stroigg, xiwjiwjd2)
	const PI = 3.14141414
	const (
		x1iwjiwjd  int = 1
		x1iwjiwjd2 int = 20
		st1roigg       = "owksowksowks"
	)
	perimeter, _ := doOperationOnRectange(10, 20)
	fmt.Println(perimeter)

	calledX := xfunc
	calledX(10, 20)

	calledY := yfunc
	calledY(10, 20)
	const X = 1212
	swjiwjs := 22222222222222
	var (
		yaafunc = func(y int) {
			fmt.Println(swjiwjs)
			fmt.Println(y)
		}
	)
	yaafunc(111111111111111)
	variadicExample2("jw2ijw2", swjiwjs)
	//We can have a serious look at it
	variadicExample2("ajiiwhsw", 873, 238738273, []int{1, 23, 4}, [][]int{{1, 2}}, []string{"jswijsiw", "isjwijsiwj"}, X)
	def()
	HighOrderFunctions()
	num1 := 10
	if num1 < 20 {
		fmt.Println("Number is less than 20")
	} else if num1 == 20 {
		fmt.Println("Number is equal to 20")
	} else {
		fmt.Println("Number is greater than 20")

	}
}
