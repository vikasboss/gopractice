package main

import "fmt"

func main() {
	var array1 [5]int
	array1[1] = 10
	for _, v := range array1 {
		fmt.Println(v)
	}

	array2 := [5]int{1, 2, 3, 4, 5}
	var y [5]int = [5]int{0, 5, 10, 15, 20}
	k := [...]int{10, 20, 30}

	// Initialize values for specific array elements
	a := [5]int{1: 1, 3: 25}
	fmt.Println(array2, y, k, a)
	x := [5]int{0, 5, 10, 15, 20}

	// Standard for loop
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	// Getting index and value using range
	for index, element := range x {
		fmt.Println(index, "=>", element)
	}

	// Only getting value using range
	for _, value := range x {
		fmt.Println(value)
	}

	// Range and counter
	j := 0
	for range x {
		fmt.Println(x[j])
		j++
	}
	xa := [5]int{0, 5, 10, 15, 20}

	// Copy array values
	y := x

	// Copy by reference
	z := &xa

	fmt.Printf("x: %v\n", xa)
	fmt.Printf("y: %v\n", y)

	xa[0] = 1

	fmt.Printf("x: %v\n", xa)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", *z)

}
