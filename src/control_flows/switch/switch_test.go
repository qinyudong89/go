package switch_test

import (
    "fmt"
    "math/rand"
    "time"
    "testing"
)

func TestSwitch(t *testing.T) {
    sec := time.Now().Unix()
    t.Log("sec: ",sec)
    rand.Seed(sec)
    i := rand.Int31n(10)
    t.Log("i: ",i)
    switch i {
    case 0:
        fmt.Print("zero...")
    case 1:
        fmt.Print("one...")
    case 2:
        fmt.Print("two...")
    default :
        fmt.Print(i)
    }
}
