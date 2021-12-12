//切片的遍历
//切片的遍历方式和数组是一致的，支持索引遍历和for range遍历
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	//for循环遍历切片
	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("index:%d value:%v\n", i, s[i])
	// 	//fmt.Println(i, s[i])
	// }

	//for range遍历切片
	for index, value := range s {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}
