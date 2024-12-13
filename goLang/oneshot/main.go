package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ab, bc := "hello", 4
	wg.Add(1)
	go send(ab, bc) // creates a lightweight thread
	abc := fmt.Sprintf("%v hello from %v", ab, bc)
	fmt.Print(abc)
	// time.Sleep(5 * time.Second) // main thread dont wait for anny other thread
	wg.Wait()
}

var wg = sync.WaitGroup{}
/*
	contains 3 functions :Add(adds count for main thraead to wait)
						  Done(reduce the waiting count)
						  Wait(make the main thread wait)
*/

func send(a string, b int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("%v %v", a, b)
	wg.Done()
}
