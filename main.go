package main

import "fmt"

func main() {
	Show("main_window")
	// TestFunc()
}

func TestFunc() {
	// 在一个函数中再次创建一个函数，要用到匿名函数。
	func() {
		fmt.Println("bbb")
	}() // 表示调用执行匿名函数
	fmt.Println("aaa")

}
