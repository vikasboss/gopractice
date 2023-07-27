package basics

import "fmt"

func main() {
	PrintHello()
}

func PrintWorld() {
	fmt.Print(" World")
}

// Defer is a spezial statement that schedules functions to be executed after the function completes

func PrintHello() {
	defer PrintWorld()
	defer fmt.Print("Hello")

}
