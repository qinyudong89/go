package main

import "fmt"

type name struct {
  width, height int
}

// 为 rect 结构体定义计算面积方法 area(), 参数为指针类型
func area(r *rect) int {
	return r.width * r.height
}
// 为 rect 结构体定义计算周长的方法 perim(), 参数为值类型
func perim(r *rect) int {
	return 2 * (r.width + r.height)
}

func main() {
    // 定义一个矩形结构体变量
	r := rect{width: 10, height: 4}
	// 分别调用计算面积和周长的方法
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

    // 指针方式调用
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
