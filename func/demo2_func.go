//可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
//注意：可变参数通常要作为函数的最后一个参数。
package main

import (
	"fmt"
)

//可变参数使用
// func main() {
// 	s1 := intSum2()
// 	s2 := intSum2(10)
// 	s3 := intSum2(10, 20, 30)
// 	fmt.Println(s1, s2, s3)
// }

// func intSum2(x ...int) int {
// 	fmt.Println(x) //x是一个切片
// 	sum := 0
// 	for _, v := range x {
// 		sum += v
// 	}
// 	return sum
// }

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
//本质上，函数的可变参数是通过切片来实现的
func main() {
	s5 := intSum3(100)
	s6 := intSum3(100, 10)
	s7 := intSum3(100, 10, 20)
	fmt.Println(s5, s6, s7)
}

func intSum3(x3 int, y3 ...int) int {
	fmt.Println(x3, y3)
	sum3 := x3
	for _, v3 := range y3 {
		sum3 += v3
	}
	return sum3
}
