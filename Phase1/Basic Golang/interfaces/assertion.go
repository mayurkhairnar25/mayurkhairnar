package main

import "fmt"

func myvalue (a interface {} )  {
	value,ok :=a.(float64)
	fmt.Println(value,ok)
}
func main ()  {
	var a1 interface{

	}=98.25
	myvalue(a1)

	var a2 interface{

	}= "Hello"
	myvalue(a2)
}
