package main

import (
	"fmt"
	"time"
)

func runner1() {
	fmt.Println("Runner 1 ")
}
func runner2() {
	fmt.Println("Runner 2")
}
func execute() {
	go runner1()
	go runner2()

}
func main() {
	execute()
	time.Sleep(100 * time.Millisecond)
}

