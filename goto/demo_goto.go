//goto语句通过标签进行代码间的无条件跳转。
//goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
//Go语言中使用goto语句能简化一些代码的实现过程
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				goto breakTag //设置退出标签
			}
			fmt.Printf("i=%d;j=%d\n", i, j)
		}
		return
	}
breakTag: //退出标签
	fmt.Println("结束for循环")
}
