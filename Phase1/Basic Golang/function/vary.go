package main

import "fmt"

func main ()  {
	ve ()
	ve ("Hello")
}
func ve (s ... string ){
	fmt.Println(s)
}