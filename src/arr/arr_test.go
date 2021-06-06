package arr_test

import "testing"

func TestArr(t *testing.T)  {
  cites := [5]string{"GuangZhou","ShenZhen","ShangHai"}
  t.Log("cites:",cites)
  t.Log(len(cites))
}
