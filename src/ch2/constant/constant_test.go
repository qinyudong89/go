package constant

import "testing"
const (
  m = iota+100
  tu
  med
)

func TestConstant(t *testing.T)  {
  t.Log(m,med)
}
