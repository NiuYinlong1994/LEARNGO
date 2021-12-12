//要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
package main

import (
	"fmt"
)

func main() {
	s1 := []string{"北京", "上海", "广州", "深圳"}
	fmt.Println(s1)
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1)
}
