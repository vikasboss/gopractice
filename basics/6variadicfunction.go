package basics

import (
	"fmt"
	"reflect"
)

func main() {
	printMultipleStrings("Hello", "World", "!")
	printMultipleVariables(1, "green", false, 1.314, []string{"foo", "bar", "baz"})

}

// passing multiple attributes to variadic function
func printMultipleStrings(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// Pass multiple different datatypes
func printMultipleVariables(multipleVariable ...interface{}) {
	for _, v := range multipleVariable {
		fmt.Println(v, "------>", reflect.ValueOf(v).Type())

	}
}
