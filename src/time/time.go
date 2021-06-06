package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("当前时间：", t.Unix())
	fmt.Println("当前时间字符串：", t.String())
	fmt.Println("YYYY-MM-DD : ", t.Format("2005-01-02"))
}
