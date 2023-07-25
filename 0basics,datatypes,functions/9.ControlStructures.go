package main

import "fmt"

func main() {
	x := true
	if x == true {
		fmt.Println("X is true")
	} else if x == false {
		fmt.Println("X is not true")
	}

	//kind of initialisation
	if a := 10; a == 10 {
		fmt.Println("A is equal to 10")
	}
}
