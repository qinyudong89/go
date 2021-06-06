package my_func_test

import (
	"testing"
)

func TestMyFunc(t *testing.T) {
	var num1 int = 10
	var num2 int = 20
	var sum int = sum(num1, num2)
	t.Log(sum)
	var perimeter, area = calc(num1, num2)
	t.Log(perimeter, area)

	firstName := "John"
	updateName(firstName)
	t.Log(firstName)

	//     & 运算符使用其后对象的地址。
	//    * 运算符取消引用指针。 也就是说，你可以前往指针中包含的地址访问其中的对象。
	updateName2(&firstName)
	t.Log(firstName)
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}

// 计算长方形的周长、面积
func calc(long int, width int) (perimeter int, area int) {
	perimeter = (long + width) * 2
	area = long * width
	return
}

func updateName(name string) {
	name = "David"
}

func updateName2(name *string) {
	*name = "David"
}
