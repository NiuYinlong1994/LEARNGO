//切片的定义
//var name []T
//name:表示变量名；T:表示切片中的元素类型
package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}
