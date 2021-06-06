package type_conversion_test

import
(
  "strconv"
  "testing"
  )

func TestTypeConversion(t *testing.T)  {
  // string convers int
    num, _ := strconv.Atoi("42")
    str := strconv.Itoa(num)
    t.Log(num, str)
}
