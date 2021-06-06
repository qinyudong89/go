package main
import "fmt"

type Student struct {
  name string
  grades []int
  age int
}
func (s *Student) setAge(age int)  {
  s.age = age
}

func (s *Student) getAvgGrade() float32{
  sum := 0
  for _, v := range s.grades {
    sum += v
  }
  return float32(sum)/float32(len(s.grades))
}

func (s *Student) getMaxGrade() int  {
  cur := 0
  for _, v := range s.grades{
    if cur < v {
      cur = v
    }
  }
  return cur

}

func main()  {
  r := rect{num:10}
  fmt.Printf("num: %d \n", r.add())
  s1 := Student{"Tom", []int{80,66,77}, 21}
  fmt.Println(s1)
  s1.setAge(34)
  // avg := s1.getAvgGrade()
  max := s1.getMaxGrade()
  fmt.Println(max)
}
