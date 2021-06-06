package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello world")
	} else {
		fmt.Println(os.Args[1])
	}
	num := os.Args[1]
	result := (20 / num)
	fmt.Printf("result: %d", result)
}
