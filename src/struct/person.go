package main

import "fmt"

type Person struct {
  name string
  age int
}

type Employee struct {
  id int
  info Person
}

func main()  {
  p := Person{name:"Tom",age:23}
  var employee Employee
  employee.id = 100
  employee.info = p
  employee.info.name = "Morgan"
  employee.info.age = 33
  fmt.Println(employee)
}
