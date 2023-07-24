package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
You are required to convert string to float
Convert String to Boolean
Float to String
Converting float to int
*/
func main() {
	aString := "abcd"
	fmt.Println(aString)
	//reflect is used to check for types in golang
	fmt.Println(reflect.TypeOf(aString))

	//You are required to convert string to float
	numString := "651.73829192"
	fmt.Println(strconv.ParseFloat(numString, 2))

	//Convert String to Boolean
	boolString := "false"
	fmt.Println(strconv.ParseBool(boolString))

	//Float to String
	var floatValue float64
	floatValue = 1233.23444
	fmt.Printf("%v %T", strconv.FormatFloat(floatValue, 'E', -1, 32), floatValue)

	//float to int
	fmt.Printf("%T %v", int(floatValue), int(floatValue))
}
