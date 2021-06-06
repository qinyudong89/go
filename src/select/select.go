package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Printf("exec: %v \n", <-ch)
	default:
		fmt.Printf("exec: %v \n", "Default case ...")
	}
}
