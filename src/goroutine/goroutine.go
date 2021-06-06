package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f1() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		fmt.Printf("A: %v \n", i)
	}
}

func f2() {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		fmt.Printf("B: %v \n", i)
	}
}

func main() {

	wg.Add(2)

	go f1()
	go f2()
	wg.Wait()
}
