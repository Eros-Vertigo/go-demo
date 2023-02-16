package closure

import (
	"fmt"
)

// 用途
// 1.保存状态
// 2.延迟执行
// 3.封装实现细节
// 4.实现回调函数
// 5.作为装饰器函数
// 6.实现私有变量和方法
// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

func init() {
	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}
