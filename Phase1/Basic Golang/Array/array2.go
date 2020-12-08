package main

import "fmt"

func main() {
	a := [...]int{10, 20, 30, 40, 50}
	fmt.Println(a)
	var snames [2] string
	snames[0] = "The"
	snames[1] ="Mayur"
	fmt.Println(snames)
	fmt.Println(snames[1])
	fmt.Println(len(a))



}
