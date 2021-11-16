package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition in golang")

	wg := &sync.WaitGroup{}

	score := []int{0}

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine no 1")
		score = append(score, 1)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine no 2")
		score = append(score, 2)
		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		fmt.Println("Goroutine no 3")
		score = append(score, 3)
		wg.Done()
	}(wg)

	wg.Wait()

	fmt.Println(score)

}
