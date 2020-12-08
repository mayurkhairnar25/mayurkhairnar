package main

import "fmt"

func main ()  {
	m:= add(10,15)
	fmt.Println("sum:",m)
}
func add (a,b int ) int {
	return a+b
}