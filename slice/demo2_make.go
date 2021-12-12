//make([]T, size, cap)
//T:切片的元素类型;size:切片中元素的数量;cap:切片的容量
//切片的本质就是对底层数组的封装，
//它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断
package main

import (
	"fmt"
)

func main() {
	a := make([]int, 2, 10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
