package main

import "fmt"
import "sort"

func main (){
	s:= []int {10,20,30,25,5}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

}