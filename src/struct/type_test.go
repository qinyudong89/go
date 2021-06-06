package type_test

import "testing"

type Employee struct {
    id  int
    name string

}

func TestType(t *testing.T){
  e := Employee{
    id:2000,
    name:"Tom",
  }
  t.Log(e.name)
}
