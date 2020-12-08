package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go sayhello()
		m.Lock()
		go increamnet()
	}
	wg.Wait()
}
func sayhello() {
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}
func increamnet() {
	counter++
	m.Unlock()
	wg.Done()
}

