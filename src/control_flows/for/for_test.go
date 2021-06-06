package for_test

import (
	"math/rand"
	"testing"
	// "time"
)

func TestBaseFor(t *testing.T) {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	t.Log("sum: ",sum)
  t.Log("------------------------------------------")
}

func TestWhile(t *testing.T) {
	var num int
	rand.Seed(2)
	for num != 5 {
		num = rand.Intn(10)
		t.Log(num)
	}
  t.Log("------------------------------------------")
}


func TestForBreak(t *testing.T){
  count := 10
  for i := 1; i < count; i++ {
    t.Log("i: ", i)
    if i % 3 == 0 {
      break
    }
  }
  t.Log("------------------------------------------")
}

func TestForContinue(t *testing.T){
  count := 10
  for i := 1; i < count; i++ {
    if i % 3 == 0 {
      t.Log("i: ", i)
      continue
    }
  }
  t.Log("------------------------------------------")
}
