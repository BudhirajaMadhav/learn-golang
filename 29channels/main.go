package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	wg := &sync.WaitGroup{}

	// 					   ↓an int buffer can come here
	cha := make(chan int, 0)

	// cha <- 20
	// fmt.Println(<-cha)

	wg.Add(1)
	go func(wg *sync.WaitGroup, cha chan int) {
		fmt.Println(<-cha)
		cha <- 89
		wg.Done()
	}(wg, cha)

	wg.Add(1)
	go func(wg *sync.WaitGroup, cha chan int) {
		cha <- 10
		fmt.Println(<-cha)
		wg.Done()
	}(wg, cha)

	defer wg.Wait()

}
