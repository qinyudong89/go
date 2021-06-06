package init_var_test
import "testing"

func TestInitVar(t *testing.T)  {
  var num int = 12
  var flag bool = true
  var name string = "Golang"
  t.Log("num",num)
  t.Log("bool:",flag)
  t.Log("name:",name)
}
