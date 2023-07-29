package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Go routines
func querySite(url string) {
	fmt.Println("Http request ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // Corrected the method name to Close()
	fmt.Println("Read Body: ", url)
	body, err := ioutil.ReadAll(response.Body) // Moved the declaration outside the if statement
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body length: ", len(body))
}

func main() {
	// Using a named function in a Goroutine
	go querySite("https://gabrieltanner.org")
	go querySite("https://stackoverflow.com")
	// Using an anonymous inner function in a goroutine
	go func(x int) {
		fmt.Println("Number: ", x)
	}(100)

	// Add some delay to allow goroutines to complete before the program exits
	// In a real application, you would use synchronization mechanisms like wait groups
	// to ensure all goroutines finish before the main function exits.
	fmt.Scanln()
}
