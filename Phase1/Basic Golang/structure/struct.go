package main

import "fmt"

type Person struct {
Name string
Age int
}
func main ()  {
	p1:= Person{
		Name: "Mayur",
		Age: 26,
	}
	fmt.Println(p1)
}
