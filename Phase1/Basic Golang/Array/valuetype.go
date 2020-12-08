package main

import "fmt"

func main ()  {
	a:= [5] int {10,20,30,40,50}
	b:=a
	b[2] = 25
	fmt.Println(a)
	fmt.Println(b)


}
