package main

import "fmt"

type User interface {
	PrintName(name string)
}

type Vehicle interface {
	Alert() string
}

// Create User type for Interface
type Usr int

type Car struct{}

// Implement interface function in type
func (usr Usr) PrintName(name string) {
	fmt.Println("User Id:\t", usr)
	fmt.Println("User Name:\t", name)
}

func (c Car) Alert() string {
	return "Hup! Hup!"
}

// Define an empty interface - Often used for functions that accepts any type
func Print(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
func main() {
	var user1 User
	user1 = Usr(1)
	user1.PrintName("Gabriel")

	c := Car{}
	fmt.Println(c.Alert())

	Print(20)
}
