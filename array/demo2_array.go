//数组的遍历
package main

import (
	"fmt"
)

func main() {
	fruitarray := [...]string{"苹果", "香蕉", "西瓜", "菠萝"}
	//for循环遍历
	for i := 0; i < len(fruitarray); i++ {
		fmt.Println(fruitarray[i])
	}
	//for range遍历
	for _, i := range fruitarray {
		fmt.Println(i)
	}
}
