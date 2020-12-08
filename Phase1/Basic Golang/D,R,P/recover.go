package main

import "fmt"

func main ()  {
	panicwithrecover ()
	fmt.Printf("I need to run this")
}
func panicwithrecover (){
	defer func (){
		if err:= recover();err!=nil {
			fmt.Println(err)
		}
	}()
	panic ("Error")
}