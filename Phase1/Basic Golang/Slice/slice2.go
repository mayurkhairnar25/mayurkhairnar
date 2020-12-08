package main

import "fmt"

func main ()  {
	a:= [] int {10,20,30,40,50}
	b:= a[:]
	fmt.Println(a)
	fmt.Println(b)
  b= a[2:4]
    fmt.Println(a)
    fmt.Println(b)
}
