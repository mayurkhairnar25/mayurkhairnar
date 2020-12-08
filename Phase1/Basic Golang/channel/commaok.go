package main

import "fmt"
import "sync"

var wg= sync.WaitGroup{}
func main (){
	ch:= make (chan int ,50)
	wg.Add(2)
	go func (ch <- chan int ){
		for {
			if i,ok := <- ch;ok {
				fmt.Println(i)
			}else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func (ch chan <- int){
		ch <-42
		ch <-15
		wg.Done()
	}(ch)
	wg.Wait()
}
