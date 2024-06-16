package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(rangeStart, rangeEnd int, numCh, resultCh chan int, rnd *rand.Rand) {
	defer close(resultCh)

	for i := 0; i < 10; i++ {
		num := rnd.Intn(rangeEnd-rangeStart) + rangeStart
		numCh <- num
	}
	close(numCh)

	min := <-resultCh
	max := <-resultCh

	fmt.Printf("Найменше число: %d\n", min)
	fmt.Printf("Найбільше число: %d\n", max)
}

func findMinMax(numCh, resultCh chan int) {
	var min, max int
	first := true
	for num := range numCh {
		if first {
			min, max = num, num
			first = false
		} else {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
	}
	resultCh <- min
	resultCh <- max
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	numCh := make(chan int)
	resultCh := make(chan int)

	go generateNumbers(1, 100, numCh, resultCh, rnd)
	go findMinMax(numCh, resultCh)

	time.Sleep(2 * time.Second)
}
