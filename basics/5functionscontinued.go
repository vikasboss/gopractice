// Go language provides a special feature known as an anonymous function. An anonymous function is a function which doesn’t contain any name. It is useful when you want to create an inline function. In Go language, an anonymous function can form a closure. An anonymous function is also known as function literal.

// Go program to illustrate how
// to create an anonymous function

/*

func(parameter_list)(return_type){
// code..

// Use return statement if return_type are given
// if return_type is not given, then do not
// use return statement
return
}()


*/

/*
Example:

// Go program to illustrate how
// to create an anonymous function
package main

import "fmt"

func main() {

	  // Anonymous function
	 func(){

	    fmt.Println("Welcome! to GeeksforGeeks")
	}()

}
Output:

Welcome! to GeeksforGeeks
Important Points:

In Go language, you are allowed to assign an anonymous function to a variable. When you assign a function to a variable, then the type of the variable is of function type and you can call that variable like a function call as shown in the below example.
Example:

// Go program to illustrate
// use of an anonymous function
package main

import "fmt"

func main() {

	  // Assigning an anonymous
	 // function to a variable
	 value := func(){
	    fmt.Println("Welcome! to GeeksforGeeks")
	}
	value()

}
Output:

Welcome! to GeeksforGeeks
You can also pass arguments in the anonymous function.
Example:

// Go program to pass arguments
// in the anonymous function
package main

import "fmt"

func main() {

	  // Passing arguments in anonymous function
	func(ele string){
	    fmt.Println(ele)
	}("GeeksforGeeks")

}
Output:

GeeksforGeeks
You can also pass an anonymous function as an argument into other function.
Example:

// Go program to pass an anonymous
// function as an argument into
// other function
package main

import "fmt"

	 // Passing anonymous function
	// as an argument
	func GFG(i func(p, q string)string){
	    fmt.Println(i ("Geeks", "for"))

	}

	func main() {
	    value:= func(p, q string) string{
	        return p + q + "Geeks"
	    }
	    GFG(value)
	}

Output:

GeeksforGeeks
You can also return an anonymous function from another function.
Example:

// Go program to illustrate
// use of anonymous function
package main

import "fmt"

	// Returning anonymous function
	func GFG() func(i, j string) string{
	    myf := func(i, j string)string{
	         return i + j + "GeeksforGeeks"
	    }
	   return myf
	}

	func main() {
	    value := GFG()
	    fmt.Println(value("Welcome ", "to "))
	}

Output:

Welcome to GeeksforGeeks
*/
package basics

import "fmt"

// Defining a anonymous function
var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	area := area(10, 10)
	fmt.Println(area)
}

/*
package main

import "fmt"

func main() {
	l := 10
	b := 10

	// Closure functions are a special case of a anonymous function where you access outside variables
	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()
}

*/
