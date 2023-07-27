package basics

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Infinite loop
	i := 0
	for {
		fmt.Println("Hello World!")
		// Breaks/Stops the infinite loop
		if i == 10 {
			break
		}
		i++
	}
	strings1 := []string{"Vikas", "Is", "A"}
	for index, val := range strings1 {
		fmt.Println(index, val)
	}
	for _, val := range strings1 {
		fmt.Println(val)
	}
	for index := range strings1 {
		fmt.Println(index)
	}
	// Basic while loop
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Do while loop
	num := 0
	for {
		// Work
		fmt.Println(num)

		if num == 10 {
			break
		}
		num++
	}

}
