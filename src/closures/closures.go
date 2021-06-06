package main

import "fmt"

func incr() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := incr()
	for i := 0; i < 10; i++ {
		fmt.Println(next())
	}

}
