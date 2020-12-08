package main

import "fmt"

type Vehicle struct {
	Name string
	Type string
}
type Car struct {
	Vehicle
	MaxSpeed int
}
func main ()  {
	C := Car {}
	C.Name = "BMW"
	C.Type = "SPORT"
	C.MaxSpeed = 250

	fmt.Println(C)
}
