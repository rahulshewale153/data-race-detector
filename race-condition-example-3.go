package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	EvenCount int
	OddCount  int
}

var c Counter

func main() {
	numberCollection := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("Start Goroutine")
	var wg sync.WaitGroup
	wg.Add(11)
	for _, number := range numberCollection {
		go setCounter(&wg, number)
	}
	wg.Wait()
	fmt.Printf("Total Event Number is %v and Odd Number is %v\n", c.EvenCount, c.OddCount)
}
func setCounter(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	if number%2 == 0 {
		c.EvenCount++
		return
	}
	c.OddCount++
}
