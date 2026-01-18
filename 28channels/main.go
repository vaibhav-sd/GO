package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO")
	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5
	wg.Add(2)

	// Read Only <-cha
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isOpen := <-myCh

		fmt.Println(isOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	// Send Only chan <-
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}
