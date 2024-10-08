package main

import (
	"fmt"
	"math/rand"
)

type Work struct {
	value int
}

type Result struct {
	value int
}

func main() {
	workCh := make(chan Work)
	resultCh := make(chan Result)
	done := make(chan bool)

	workQuene := make([]Work, 100)
	for i := range workQuene {
		workQuene[i].value = rand.Int()
	}
	// Create 10 worker goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for work := range workCh {
				// Compute result
				result := Result{
					value: work.value * 2,
				}
				// Send the result via the result channel
				resultCh <- result
			}
		}()
	}
	results := make([]Result, 0)
	go func() {
		// Collect all the results
		for i := 0; i < len(workQuene); i++ {
			results = append(results, <-resultCh)
		}
		// When all the results are collected, notify the done channel
		done <- true
	}()
	// Send all the work to the workers
	for _, work := range workQuene {
		workCh <- work
	}
	close(workCh)
	// Wait until everything is done
	<-done
	fmt.Println(results)
}
