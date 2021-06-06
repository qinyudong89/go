package main
import "fmt"

type user struct {
  id int
  name string
  age int
}
func main()  {
  u := user{id:1001,name:"Tom",age:25}
  u.age=90
  fmt.Println(u)
}
