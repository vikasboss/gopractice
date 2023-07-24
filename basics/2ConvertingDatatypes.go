/*
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     IEEE-754 32-bit floating-point numbers
float64     IEEE-754 64-bit floating-point numbers
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

uint     unsigned, either 32 or 64 bits
int      signed, either 32 or 64 bits
uintptr  unsigned integer large enough to store the uninterpreted bits of a pointer value
*/
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
	// Declaring multiple variables
	firstName, lastName := "FirstName", "LastName"

	fmt.Println(firstName + lastName)

	// Variable Declaration Block
	var (
		name = "Donald Duck"
		age  = 50
	)

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
