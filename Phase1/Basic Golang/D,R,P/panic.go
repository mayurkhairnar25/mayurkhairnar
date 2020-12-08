package main

import "fmt"

func entry (large *string , Aname *string ){
	if large ==nil {
		panic ("error ")
	}
	if Aname == nil {
		panic ("error")
	}
	fmt.Printf("%s,%s",*large,*Aname)
}
func main ()  {
	Alarge := "Go"
	defer fmt.Println("Defer ")
	entry(&Alarge,nil)
}