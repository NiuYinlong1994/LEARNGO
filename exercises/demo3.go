//编写代码打印9*9乘法表
package main

import (
	"fmt"
)

func main() {
	var multi int
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			multi = j * i
			fmt.Printf("%v*%v=%v\t", i, j, multi)
		}
		fmt.Println()
	}
}
