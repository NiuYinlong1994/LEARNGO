//拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
package main

import (
	"fmt"
)

func main() {
	//由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化
	// a := make([]int, 3) //[0 0 0]
	// b := a              //将a直接赋值给b,a和b共用一个底层数组
	// b[0] = 100
	// fmt.Println(a)
	// fmt.Println(b)

	//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中
	//copy(destSlice, srcSlice []T)
	//srcSlice: 数据来源切片; destSlice: 目标切片
	x := []int{1, 2, 3, 4, 5}
	y := make([]int, 5, 5)
	copy(y, x) //使用copy()函数将切片x中的元素复制给切片y
	fmt.Println(x)
	fmt.Println(y)
	y[0] = 1000
	fmt.Println(x)
	fmt.Println(y)
}
