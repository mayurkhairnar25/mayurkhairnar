package main

import "fmt"
import "errors"

func main() {

	// create error
	myErr := errors.New("Something unexpected happend!")

	// print error message
	fmt.Println(myErr)
}

