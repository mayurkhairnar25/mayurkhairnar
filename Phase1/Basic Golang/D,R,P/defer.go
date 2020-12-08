package main

import "fmt"

func main ()  {
	fmt.Println("One")
	defer fmt.Println("Two")
	defer  fmt.Println("Three")
}
