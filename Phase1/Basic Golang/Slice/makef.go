package main

import "fmt"

func main ()  {
	s:= make ([] int ,3)
	fmt.Println(s)
	s=append(s,10,20)
	fmt.Println(s)
	s=append(s,30,40,50)
	fmt.Println(s)

	s=append(s,[] int {10,20}...)
	fmt.Println(s)

}
