package basics

import (
	"fmt"
)

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("This is 1")
	case 2, 3:
		fmt.Println("This is 2 or 3")
	default:
		fmt.Println("This is default case")
	}

	dayofweek := 4
	switch dayofweek {
	case 1:
		fmt.Println("Go to Work")
		fallthrough
	case 2:
		fmt.Println("Buy some bread")
		fallthrough
	case 3:
		fmt.Println("Visit a Friend")
	default:
		fmt.Println("This is the default case")
	}
	// Switch using conditional statements
	switch {
	case num < 5:
		fmt.Println("Smaller than 5")
	case num == 5:
		fmt.Println("Five")
	case num > 5:
		fmt.Println("Bigger than five")
	default:
		fmt.Println("No information about the number")
	}

	// Switch initializer statement
	switch number := 5; {
	case number < 5:
		fmt.Println("Smaller than 5")
	case number == 5:
		fmt.Println("Five")
	case number > 5:
		fmt.Println("Bigger than five")
	default:
		fmt.Println("No information about the number")
	}

}
