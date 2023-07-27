package main

import (
	"fmt"
	"os"
)

func printExamples() {
	// fmt.Printf is used for formatted printing to the standard output. It takes a format string with verbs (format specifiers) that define how to format the provided values. It works similarly to C's printf function. The formatted string is printed to the standard output without adding a newline.
	name := "Alice"
	age := 28
	fmt.Printf("My name is %s and I am %d years old.\n", name, age)
	//	fmt.Sprintf is used for formatted string creation, similar to fmt.Printf. It returns the formatted string but does not print it to the standard output. This function is helpful when you want to store the formatted string in a variable.
	name = "Bob"
	age = 35
	formattedStr := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedStr)
	//fmt.Fprintf is used for formatted printing to a specified io.Writer (e.g., file, network connection) instead of the standard output. It takes the writer as the first argument, followed by the format string and values to format. It's similar to fmt.Printf, but it allows you to print to a custom writer.
	name = "Charlie"
	age = 42
	file, _ := os.Create("output.txt")
	defer file.Close()

	fmt.Fprintf(file, "My name is %s and I am %d years old.\n", name, age)

	//fmt.Print is a function used to print text and values to the standard output (usually the console or terminal). It takes a list of arguments and prints them as they are, without any formatting. The values are separated by a space, and there is no newline added at the end.
	name = "John"
	age = 30
	fmt.Print("My name is", name, "and I am", age, "years old.")

}

func main() {
	printExamples()
}
