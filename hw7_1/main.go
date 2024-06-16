package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan int, rnd *rand.Rand) {
	for i := 0; i < 10; i++ {
		num := rnd.Intn(100) 
		ch <- num
	}
	close(ch)
}

func calculateAverage(inputCh, outputCh chan int) {
	var sum, count int
	for num := range inputCh {
		sum += num
		count++
	}
	average := sum / count
	outputCh <- average
	close(outputCh)
}

func printAverage(ch chan int) {
	avg := <-ch
	fmt.Printf("Середнє значення: %d\n", avg)
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	numCh := make(chan int)
	avgCh := make(chan int)

	go generateNumbers(numCh, rnd)
	go calculateAverage(numCh, avgCh)
	go printAverage(avgCh)

	time.Sleep(2 * time.Second)
}
