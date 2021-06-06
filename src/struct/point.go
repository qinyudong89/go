package main
import "fmt"
type Point struct {
  x int
  y int
}

func changeX(pt Point) {
  pt.x=100
}

func main()  {
  p := Point{x:2,y:3}
  p2 := &p
  fmt.Println(p2)
  p.x = 100
  fmt.Println(p,p2)
}
