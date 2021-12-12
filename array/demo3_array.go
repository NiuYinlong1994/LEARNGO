//二维数组的定义以及遍历
package main

import (
	"fmt"
)

func main() {
	//二维数组的定义:
	// ballArray := [3][2]string{
	// 	{"足球", "排球"},
	// 	{"篮球", "羽毛球"},
	// 	{"乒乓球", "保龄球"},
	// }
	// fmt.Println(ballArray)
	// fmt.Printf("type of ballArray:%T\n", ballArray)
	// //索引找到羽毛球
	// fmt.Println(ballArray[1][1])

	//二维数组的遍历
	nameArray := [...][3]string{
		{"小红", "小黄", "小蓝"},
		{"大黑", "大白", "大紫"},
		{"阿猫", "阿狗", "阿虎"},
	}
	for _, i := range nameArray {
		for _, j := range i {
			fmt.Printf("%s\t", j)
		}
		fmt.Println()
	}
}
