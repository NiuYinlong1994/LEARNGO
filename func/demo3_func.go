//Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println(calc(20, 10))
// }
// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	sub := x - y
// 	return sum, sub
// }

//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
func main() {
	fmt.Println(calc(20, 10))
}
func calc(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

//当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片
// func someFunc(x string) []int {
// 	if x == "" {
// 		return nil		//没有必要返回[]int{}
// 	}
//	...
// }
