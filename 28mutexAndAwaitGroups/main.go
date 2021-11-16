package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in golang")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	score := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine no 1")

		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine no 2")

		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("Goroutine no 3")

		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		
		wg.Done()
	}(wg, mut)

	wg.Wait()

	fmt.Println(score)

}
