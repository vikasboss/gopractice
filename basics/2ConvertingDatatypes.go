package main

import (
	"fmt"
	"reflect"
)

func main() {
	aString := "abcd"
	fmt.Println(aString)
	//reflect is
	fmt.Println(reflect.TypeOf(aString))

}
