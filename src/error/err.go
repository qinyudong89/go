package main

import "fmt"

func Hello(ch chan int) {
	fmt.Println("hello everybody , I'm asong")
	ch <- 1
}

func main() {
	ch := make(chan int)
	go Hello(ch)
	str := <-ch
	fmt.Println("Golang梦工厂=====>", str)

}
