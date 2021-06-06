package if_else_test

import "testing"

func TestIfElse(t *testing.T)  {
  score := 60
  if score < 60 {
    t.Log("score: 不及格")
  }else if score >= 60 && score < 80 {
    t.Log("良好")
  }else {
    t.Log("优秀")
  }
}
