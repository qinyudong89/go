package main

import "fmt"

func send(ch chan string, msg string) {
	ch <- msg
}

func main() {
	size := 10
	ch := make(chan string, size)
	go send(ch, "one")
	go send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}
