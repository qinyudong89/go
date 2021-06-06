package main

import (
	"fmt"
	"os"
)

func mian() {
	filename := "‪C:\\Users\\qinyu\\Desktop\\tmp.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error ")
	}
	/**
	* 这里defer要写在err判断的后边而不是os.Open后边
	* 如果资源没有获取成功，就没有必要对资源执行释放操作
	* 如果err不为nil而执行资源执行释放操作，有可能导致panic
	 */
	defer f.Close()

}
