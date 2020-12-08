package main

import "fmt"

func main ()  {
	for i:=0;i<5;i++ {
		sayhello("Hello Go",i)
	}
}
func sayhello (msg string , idx int ){
	fmt.Println(msg,idx)
}